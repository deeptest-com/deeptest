package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.IssuePrioritySolutionDao;
import com.ngtesting.platform.model.IsuPrioritySolution;
import com.ngtesting.platform.service.IssuePrioritySolutionService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class IssuePrioritySolutionServiceImpl extends BaseServiceImpl implements IssuePrioritySolutionService {

	@Autowired
	IssuePrioritySolutionDao solutionDao;

	@Override
	public List<IsuPrioritySolution> list(Integer orgId) {
		List<IsuPrioritySolution> ls = solutionDao.list(orgId);

		return ls;
	}

    @Override
    public List<IsuPrioritySolution> list(Integer orgId, Integer prjId) {
        List<IsuPrioritySolution> ls = solutionDao.list(orgId);

        return ls;
    }

	@Override
	public IsuPrioritySolution get(Integer id, Integer orgId) {
		return solutionDao.get(id, orgId);
	}

	@Override
	public IsuPrioritySolution save(IsuPrioritySolution vo, Integer orgId) {

		if (vo.getId() == null) {
			vo.setOrgId(orgId);
			solutionDao.save(vo);
		} else {
			Integer count = solutionDao.update(vo);
			if (count == 0) {
				return null;
			}
		}

		return vo;
	}

	@Override
	public Boolean delete(Integer id, Integer orgId) {
		Integer count = solutionDao.delete(id, orgId);
		if (count == 0) {
			return false;
		}

		return true;
	}

}
