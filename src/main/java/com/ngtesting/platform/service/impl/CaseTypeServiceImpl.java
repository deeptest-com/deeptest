package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.CaseTypeDao;
import com.ngtesting.platform.model.TstCaseType;
import com.ngtesting.platform.service.CaseTypeService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@Service
public class CaseTypeServiceImpl extends BaseServiceImpl implements CaseTypeService {
    @Autowired
	CaseTypeDao caseTypeDao;

	@Override
	public List<TstCaseType> list(Integer orgId) {
        List<TstCaseType> ls = caseTypeDao.list(orgId);

		return ls;
	}

    @Override
    public TstCaseType get(Integer id, Integer orgId) {
        return caseTypeDao.get(id, orgId);
    }

    @Override
	public TstCaseType save(TstCaseType vo, Integer orgId) {
        vo.setOrgId(orgId);

        if (vo.getId() == null) {
            Integer maxOrder = caseTypeDao.getMaxOrdrNumb(orgId);
            if (maxOrder == null) {
                maxOrder = 0;
            }
            vo.setOrdr(maxOrder + 10);

            caseTypeDao.save(vo);
        } else {
            caseTypeDao.update(vo);
        }

        return vo;
	}

	@Override
	public Boolean delete(Integer id, Integer orgId) {
        caseTypeDao.delete(id, orgId);

		return true;
	}

	@Override
    @Transactional
	public Boolean setDefault(Integer id, Integer orgId) {
        caseTypeDao.removeDefault(orgId);
        caseTypeDao.setDefault(id, orgId);

		return true;
	}

	@Override
    @Transactional
	public Boolean changeOrder(Integer id, String act, Integer orgId) {
        TstCaseType curr = caseTypeDao.get(id, orgId);
        if (curr == null) {
            return false;
        }

        TstCaseType neighbor = null;
        if ("up".equals(act)) {
            neighbor = caseTypeDao.getPrev(curr.getOrdr(), orgId);
        } else if ("down".equals(act)) {
            neighbor = caseTypeDao.getNext(curr.getOrdr(), orgId);
        }
        if (neighbor == null) {
            return false;
        }

        Integer currOrder = curr.getOrdr();
        Integer neighborOrder = neighbor.getOrdr();
        caseTypeDao.setOrder(id, neighborOrder, orgId);
        caseTypeDao.setOrder(neighbor.getId(), currOrder, orgId);

        return true;
	}

}
