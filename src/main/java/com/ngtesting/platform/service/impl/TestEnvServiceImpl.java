package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.dao.TestEnvDao;
import com.ngtesting.platform.model.TstEnv;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.TestEnvService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@Service
public class TestEnvServiceImpl extends BaseServiceImpl implements TestEnvService {

    @Autowired
    TestEnvDao envDao;

    @Override
    public List<TstEnv> list(Integer projectId, String keywords, Boolean disabled) {
        List<TstEnv> ls = envDao.query(projectId, keywords, disabled);
        return ls;
    }

    @Override
    public TstEnv getById(Integer id, Integer projectId) {
        TstEnv po = envDao.get(id, projectId);
        return po;
    }

    @Override
    @Transactional
    public TstEnv save(JSONObject json, TstUser user) {
        TstEnv vo = JSON.parseObject(JSON.toJSONString(json), TstEnv.class);
        Integer id = vo.getId();

        if (id == null) {
            vo.setProjectId(user.getDefaultPrjId());
            vo.setOrgId(user.getDefaultOrgId());

            Integer maxOrder = envDao.getMaxOrdrNumb(vo.getProjectId());
            if (maxOrder == null) {
                maxOrder = 0;
            }
            vo.setOrdr(maxOrder + 10);

            envDao.add(vo);
        } else {
            Integer count = envDao.update(vo);
            if (count == 0) {
                return null;
            }
        }

        return vo;
    }

    @Override
    @Transactional
    public Boolean delete(Integer id, Integer projectId) {
        Integer count = envDao.delete(id, projectId);
        if (count == 0) {
            return false;
        }

        return true;
    }

    @Override
    @Transactional
    public Boolean changeOrder(Integer id, String act, Integer projectId) {
        TstEnv curr = envDao.get(id, projectId);
        if (curr == null) {
            return false;
        }

        TstEnv neighbor = null;
        if ("up".equals(act)) {
            neighbor = envDao.getPrev(curr.getOrdr(), projectId);
        } else if ("down".equals(act)) {
            neighbor = envDao.getNext(curr.getOrdr(), projectId);
        }
        if (neighbor != null) {
            Integer currOrder = curr.getOrdr();
            Integer neighborOrder = neighbor.getOrdr();
            envDao.setOrder(id, neighborOrder, projectId);
            envDao.setOrder(neighbor.getId(), currOrder, projectId);
        }

        return true;
    }

}

