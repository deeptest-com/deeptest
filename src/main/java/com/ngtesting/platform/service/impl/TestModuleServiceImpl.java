package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.dao.TestModuleDao;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.model.TstModule;
import com.ngtesting.platform.service.intf.TestModuleService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@Service
public class TestModuleServiceImpl extends BaseServiceImpl implements TestModuleService {
    @Autowired
    TestModuleDao moduleDao;

    @Override
    public List<TstModule> list(Integer projectId, String keywords, Boolean disabled) {
        List<TstModule> ls = moduleDao.query(projectId, keywords, disabled);
        return ls;
    }

    @Override
    public TstModule getById(Integer id, Integer projectId) {
        TstModule po = moduleDao.get(id, projectId);
        return po;
    }

    @Override
    public TstModule save(JSONObject json, TstUser user) {
        TstModule vo = JSON.parseObject(JSON.toJSONString(json), TstModule.class);
        Integer id = vo.getId();

        if (id == null) {
            vo.setProjectId(user.getDefaultPrjId());

            Integer maxOrder = moduleDao.getMaxOrdrNumb(vo.getProjectId());
            if (maxOrder == null) {
                maxOrder = 0;
            }
            vo.setOrdr(maxOrder + 10);

            moduleDao.save(vo);
        } else {
            Integer count = moduleDao.update(vo);
            if (count == 0) {
                return null;
            }
        }

        return vo;
    }

    @Override
    public Boolean delete(Integer id, Integer projectId) {
        Integer count = moduleDao.delete(id, projectId);
        if (count == 0) {
            return false;
        }

        return true;
    }

    @Override
    @Transactional
    public Boolean changeOrder(Integer id, String act, Integer projectId) {
        TstModule curr = moduleDao.get(id, projectId);
        if (curr == null) {
            return false;
        }

        TstModule neighbor = null;
        if ("up".equals(act)) {
            neighbor = moduleDao.getPrev(curr.getOrdr(), projectId);
        } else if ("down".equals(act)) {
            neighbor = moduleDao.getNext(curr.getOrdr(), projectId);
        }
        if (neighbor != null) {
            Integer currOrder = curr.getOrdr();
            Integer neighborOrder = neighbor.getOrdr();
            moduleDao.setOrder(id, neighborOrder, projectId);
            moduleDao.setOrder(neighbor.getId(), currOrder, projectId);
        }

        return true;
    }

}

