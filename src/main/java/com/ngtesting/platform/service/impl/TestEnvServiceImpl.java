package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.TestEnvDao;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.model.TstEnv;
import com.ngtesting.platform.service.TestEnvService;
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
    public TstEnv getById(Integer id) {
        TstEnv po = envDao.get(id);
        return po;
    }

    @Override
    public TstEnv save(JSONObject json, TstUser optUser) {
        Integer id = json.getInteger("id");

        TstEnv po = null;
        TstEnv vo = JSON.parseObject(JSON.toJSONString(json), TstEnv.class);

        Constant.MsgType action;
        if (id != null) {
            action = Constant.MsgType.update;

            envDao.update(vo);
        } else {
            action = Constant.MsgType.create;

            Integer maxOrder = envDao.getMaxOrdrNumb(vo.getProjectId());
            if (maxOrder == null) {
                maxOrder = 0;
            }
            vo.setOrdr(maxOrder + 10);

            envDao.add(vo);
        }

        return vo;
    }

    @Override
    public void delete(Integer id, Integer clientId) {
        envDao.delete(id);
    }

    @Override
    @Transactional
    public boolean changeOrder(Integer id, String act, Integer projectId) {
        TstEnv curr = envDao.get(id);
        TstEnv neighbor = null;
        if ("up".equals(act)) {
            neighbor = envDao.getPrev(curr.getOrdr(), projectId);
        } else if ("down".equals(act)) {
            neighbor = envDao.getNext(curr.getOrdr(), projectId);
        }
        if (neighbor == null) {
            return false;
        }

        Integer currOrder = curr.getOrdr();
        Integer neighborOrder = neighbor.getOrdr();
        envDao.setOrder(id, neighborOrder);
        envDao.setOrder(neighbor.getId(), currOrder);

        return true;
    }

}

