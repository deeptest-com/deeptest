package com.ngtesting.platform.service.impl;

import com.github.pagehelper.PageHelper;
import com.ngtesting.platform.dao.OrgGroupDao;
import com.ngtesting.platform.model.TstOrgGroup;
import com.ngtesting.platform.service.OrgGroupService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class OrgGroupServiceImpl extends BaseServiceImpl implements OrgGroupService {

	@Autowired
	private OrgGroupDao groupDao;

	@Override
	public List<TstOrgGroup> listByPage(Integer orgId, String keywords, Boolean disabled, Integer pageNum, Integer pageSize) {
        List<TstOrgGroup> groups = groupDao.query(orgId, keywords, disabled);

        return groups;
	}

	@Override
	public List<TstOrgGroup> search(Integer orgId, String keywords, String exceptIds) {
        PageHelper.startPage(0, 20);
        List<TstOrgGroup> groups = groupDao.search(orgId, keywords, exceptIds);

		return groups;
	}

    @Override
    public List<TstOrgGroup> list(Integer orgId) {
        List<TstOrgGroup> groups = groupDao.list(orgId);

        return groups;
    }

    @Override
    public TstOrgGroup get(Integer id, Integer orgId) {
        TstOrgGroup group = groupDao.get(id, orgId);
        return group;
    }

	@Override
	public TstOrgGroup save(TstOrgGroup vo, Integer orgId) {
        vo.setOrgId(orgId);

		if (vo.getId() == null) {
			groupDao.save(vo);
		} else {
            Integer count = groupDao.update(vo);
            if (count == 0) {
                return null;
            }
        }

		return vo;
	}

	@Override
	public boolean delete(Integer id, Integer orgId) {
        Integer count = groupDao.delete(id, orgId);

		return count > 0;
	}
}
