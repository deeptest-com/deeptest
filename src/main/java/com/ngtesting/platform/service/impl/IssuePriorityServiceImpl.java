package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.IssuePriorityDao;
import com.ngtesting.platform.model.IsuPriority;
import com.ngtesting.platform.service.IssuePriorityService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@Service
public class IssuePriorityServiceImpl extends BaseServiceImpl implements IssuePriorityService {

	@Autowired
    IssuePriorityDao issuePriorityDao;

	@Override
	public List<IsuPriority> list(Integer orgId) {
		List<IsuPriority> ls = issuePriorityDao.list(orgId);

		return ls;
	}

    @Override
    public List<IsuPriority> list(Integer orgId, Integer prjId) {
        List<IsuPriority> ls = issuePriorityDao.list(orgId);

        return ls;
    }

	@Override
	public IsuPriority get(Integer id, Integer orgId) {
		return issuePriorityDao.get(id, orgId);
	}

	@Override
	public IsuPriority save(IsuPriority vo, Integer orgId) {

		if (vo.getId() == null) {
			Integer maxOrder = issuePriorityDao.getMaxOrdrNumb(orgId);
			if (maxOrder == null) {
				maxOrder = 0;
			}
			vo.setOrdr(maxOrder + 10);

			vo.setOrgId(orgId);
			issuePriorityDao.save(vo);
		} else {
			Integer count = issuePriorityDao.update(vo);
			if (count == 0) {
				return null;
			}
		}

		return vo;
	}

	@Override
	public Boolean delete(Integer id, Integer orgId) {
		Integer count = issuePriorityDao.delete(id, orgId);
		if (count == 0) {
			return false;
		}

		return true;
	}

	@Override
	@Transactional
	public Boolean setDefault(Integer id, Integer orgId) {
		issuePriorityDao.removeDefault(orgId);

		Integer count = issuePriorityDao.setDefault(id, orgId);
		if (count == 0) {
			return false;
		}
		return true;
	}

	@Override
	@Transactional
	public Boolean changeOrder(Integer id, String act, Integer orgId) {
		IsuPriority curr = issuePriorityDao.get(id, orgId);
		if (curr == null) {
			return false;
		}

		IsuPriority neighbor = null;
		if ("up".equals(act)) {
			neighbor = issuePriorityDao.getPrev(curr.getOrdr(), orgId);
		} else if ("down".equals(act)) {
			neighbor = issuePriorityDao.getNext(curr.getOrdr(), orgId);
		}
		if (neighbor == null) {
			return false;
		}

		Integer currOrder = curr.getOrdr();
		Integer neighborOrder = neighbor.getOrdr();
		issuePriorityDao.setOrder(id, neighborOrder, orgId);
		issuePriorityDao.setOrder(neighbor.getId(), currOrder, orgId);

		return true;
	}
}
