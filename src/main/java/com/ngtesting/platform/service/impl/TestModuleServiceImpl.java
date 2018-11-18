package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.dao.TestModuleDao;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.model.TstModule;
import com.ngtesting.platform.service.TestModuleService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@Service
public class TestModuleServiceImpl extends BaseServiceImpl implements TestModuleService {
    @Autowired
    TestModuleDao verDao;

    @Override
    public List<TstModule> list(Integer projectId, String keywords, Boolean disabled) {
        List<TstModule> ls = verDao.query(projectId, keywords, disabled);
        return ls;
    }

    @Override
    public TstModule getById(Integer id, Integer projectId) {
        TstModule po = verDao.get(id, projectId);
        return po;
    }

    @Override
    public TstModule save(JSONObject json, TstUser user) {
        TstModule vo = JSON.parseObject(JSON.toJSONString(json), TstModule.class);
        Integer id = vo.getId();

        vo.setProjectId(user.getDefaultPrjId());

        if (id == null) {
            Integer maxOrder = verDao.getMaxOrdrNumb(vo.getProjectId());
            if (maxOrder == null) {
                maxOrder = 0;
            }
            vo.setOrdr(maxOrder + 10);

            verDao.add(vo);
        } else {
            Integer count = verDao.update(vo);
            if (count == 0) {
                return null;
            }
        }

        return vo;
    }

    @Override
    public Boolean delete(Integer id, Integer projectId) {
        Integer count = verDao.delete(id, projectId);
        if (count == 0) {
            return false;
        }

        return true;
    }

    @Override
    @Transactional
    public Boolean changeOrder(Integer id, String act, Integer projectId) {
        TstModule curr = verDao.get(id, projectId);
        if (curr == null) {
            return false;
        }

        TstModule neighbor = null;
        if ("up".equals(act)) {
            neighbor = verDao.getPrev(curr.getOrdr(), projectId);
        } else if ("down".equals(act)) {
            neighbor = verDao.getNext(curr.getOrdr(), projectId);
        }
        if (neighbor != null) {
            Integer currOrder = curr.getOrdr();
            Integer neighborOrder = neighbor.getOrdr();
            verDao.setOrder(id, neighborOrder, projectId);
            verDao.setOrder(neighbor.getId(), currOrder, projectId);
        }

        return true;
    }

}

