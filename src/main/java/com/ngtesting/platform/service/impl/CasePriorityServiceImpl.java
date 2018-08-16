package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.CasePriorityDao;
import com.ngtesting.platform.model.TstCasePriority;
import com.ngtesting.platform.service.CasePriorityService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@Service
public class CasePriorityServiceImpl extends BaseServiceImpl implements CasePriorityService {
    @Autowired
	CasePriorityDao casePriorityDao;

	@Override
	public List<TstCasePriority> list(Integer orgId) {
		List<TstCasePriority> ls = casePriorityDao.list(orgId);

		return ls;
	}

    @Override
    public TstCasePriority get(Integer id, Integer orgId) {
        return casePriorityDao.get(id, orgId);
    }

    @Override
    @Transactional
	public TstCasePriority save(TstCasePriority vo, Integer orgId) {
        vo.setOrgId(orgId);

        if (vo.getId() == null) {
            Integer maxOrder = casePriorityDao.getMaxOrdrNumb(orgId);
            if (maxOrder == null) {
                maxOrder = 0;
            }
            vo.setOrdr(maxOrder + 10);

            casePriorityDao.save(vo);
        } else {
            Integer count = casePriorityDao.update(vo);
            if (count == 0) {
                return null;
            }
        }

        return vo;
	}

	@Override
    @Transactional
	public Boolean delete(Integer id, Integer orgId) {
        Integer count = casePriorityDao.delete(id, orgId);
        if (count == 0) {
            return false;
        }

        return true;
	}

	@Override
    @Transactional
	public Boolean setDefault(Integer id, Integer orgId) {
        casePriorityDao.removeDefault(orgId);

        Integer count = casePriorityDao.setDefault(id, orgId);
        if (count == 0) {
            return false;
        }
        return true;
	}

	@Override
    @Transactional
	public Boolean changeOrder(Integer id, String act, Integer orgId) {
        TstCasePriority curr = casePriorityDao.get(id, orgId);
        if (curr == null) {
            return false;
        }

        TstCasePriority neighbor = null;
        if ("up".equals(act)) {
            neighbor = casePriorityDao.getPrev(curr.getOrdr(), orgId);
        } else if ("down".equals(act)) {
            neighbor = casePriorityDao.getNext(curr.getOrdr(), orgId);
        }
        if (neighbor == null) {
            return false;
        }

        Integer currOrder = curr.getOrdr();
        Integer neighborOrder = neighbor.getOrdr();
        casePriorityDao.setOrder(id, neighborOrder, orgId);
        casePriorityDao.setOrder(neighbor.getId(), currOrder, orgId);

        return true;
	}

}
