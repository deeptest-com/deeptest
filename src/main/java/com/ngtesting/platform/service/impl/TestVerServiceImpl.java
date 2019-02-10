package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.dao.TestVerDao;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.model.TstVer;
import com.ngtesting.platform.service.intf.TestVerService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@Service
public class TestVerServiceImpl extends BaseServiceImpl implements TestVerService {
    @Autowired
    TestVerDao verDao;

    @Override
    public List<TstVer> list(Integer projectId, String keywords, Boolean disabled) {
        List<TstVer> ls = verDao.query(projectId, keywords, disabled);
        return ls;
    }

    @Override
    public TstVer getById(Integer id, Integer projectId) {
        TstVer po = verDao.get(id, projectId);
        return po;
    }

    @Override
    public TstVer save(JSONObject json, TstUser user) {
        TstVer vo = JSON.parseObject(JSON.toJSONString(json), TstVer.class);
        Integer id = vo.getId();

        if (id == null) {
            vo.setProjectId(user.getDefaultPrjId());
            vo.setOrgId(user.getDefaultOrgId());

            Integer maxOrder = verDao.getMaxOrdrNumb(vo.getProjectId());
            if (maxOrder == null) {
                maxOrder = 0;
            }
            vo.setOrdr(maxOrder + 10);

            verDao.save(vo);
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
        TstVer curr = verDao.get(id, projectId);
        if (curr == null) {
            return false;
        }

        TstVer neighbor = null;
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

