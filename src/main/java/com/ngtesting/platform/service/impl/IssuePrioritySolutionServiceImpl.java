package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.IssuePriorityDao;
import com.ngtesting.platform.dao.IssuePrioritySolutionDao;
import com.ngtesting.platform.model.IsuPriority;
import com.ngtesting.platform.model.IsuPrioritySolution;
import com.ngtesting.platform.service.intf.IssuePrioritySolutionService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class IssuePrioritySolutionServiceImpl extends BaseServiceImpl implements IssuePrioritySolutionService {

	@Autowired
	IssuePrioritySolutionDao solutionDao;

    @Autowired
    IssuePriorityDao priorityDao;

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

//	@Override
//	public IsuPrioritySolution get(Integer id, Integer orgId) {
//		return solutionDao.get(id, orgId);
//	}
    @Override
    public IsuPrioritySolution getDetail(Integer id, Integer orgId) {
        return solutionDao.getDetail(id, orgId);
    }

    @Override
	public IsuPrioritySolution save(IsuPrioritySolution vo, Integer orgId) {
		vo.setOrgId(orgId);

		if (vo.getId() == null) {
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

    @Override
    public Boolean setDefault(Integer id, Integer orgId) {
        solutionDao.removeDefault(orgId);

        Integer count = solutionDao.setDefault(id, orgId);
        return count > 0;
    }

	@Override
	public Boolean addPriority(Integer priorityId, Integer solutionId, Integer orgId) {
		Integer count = solutionDao.addPriority(priorityId, solutionId, orgId);
		if (count == 0) {
			return false;
		}

		return true;
	}

	@Override
	public Boolean removePriority(Integer priorityId, Integer solutionId, Integer orgId) {
		Integer count = solutionDao.removePriority(priorityId, solutionId, orgId);
		if (count == 0) {
			return false;
		}

		return true;
	}

	@Override
	public Boolean addAll(Integer solutionId, Integer orgId) {
		List<IsuPriority> priorities = priorityDao.listNotInSolution(solutionId, orgId);

		Integer count = solutionDao.addAll(priorities, solutionId, orgId);
		if (count == 0) {
			return false;
		}

		return true;
	}

	@Override
	public Boolean removeAll(Integer solutionId, Integer orgId) {
		Integer count = solutionDao.removeAll(solutionId, orgId);
		if (count == 0) {
			return false;
		}

		return true;
	}

	// For Project
	@Override
	public IsuPrioritySolution getByProject(Integer projectId, Integer orgId) {
		return solutionDao.getByProject(projectId, orgId);
	}

	@Override
	public void setByProject(Integer solutionId, Integer projectId, Integer orgId) {
		solutionDao.setByProject(solutionId, projectId, orgId);
	}

}
