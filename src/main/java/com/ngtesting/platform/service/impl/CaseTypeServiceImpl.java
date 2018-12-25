package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.CaseTypeDao;
import com.ngtesting.platform.model.TstCaseType;
import com.ngtesting.platform.service.intf.CaseTypeService;
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

        if (vo.getId() == null) {
            Integer maxOrder = caseTypeDao.getMaxOrdrNumb(orgId);
            if (maxOrder == null) {
                maxOrder = 0;
            }
            vo.setOrdr(maxOrder + 10);

            vo.setOrgId(orgId);
            caseTypeDao.save(vo);
        } else {
            Integer count = caseTypeDao.update(vo);
            if (count == 0) {
                return null;
            }
        }

        return vo;
	}

	@Override
	public Boolean delete(Integer id, Integer orgId) {
        Integer count = caseTypeDao.delete(id, orgId);

        return count > 0;
	}

	@Override
    @Transactional
	public Boolean setDefault(Integer id, Integer orgId) {
        caseTypeDao.removeDefault(orgId);

        Integer count = caseTypeDao.setDefault(id, orgId);
        return count > 0;
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
