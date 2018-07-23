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
	public List<TstOrgGroup> listByPage(Integer orgId, String keywords, String disabled, Integer pageNum, Integer pageSize) {
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
    public TstOrgGroup get(Integer id) {
        TstOrgGroup group = groupDao.get(id);
        return group;
    }

	@Override
	public TstOrgGroup save(TstOrgGroup vo, Integer orgId) {
        vo.setOrgId(orgId);

		if (vo.getId() == null) {
			groupDao.save(vo);
		} else {
            groupDao.update(vo);
        }

		return vo;
	}

	@Override
	public boolean delete(Integer id) {
//		TstOrgGroup po = (TstOrgGroup) get(TstOrgGroup.class, id);
//		po.setDeleted(true);
//		saveOrUpdate(po);

		return true;
	}

//	@Override
//	public void initDefaultBasicDataPers(TestOrg org) {
//		String [] groups = new String[]{"测试主管","测试设计","测试执行"};
//		for(String name : groups) {
//			TstOrgGroup po = new TstOrgGroup();
//			po.setName(name);
//            po.setOrgId(org.getId());
//			saveOrUpdate(po);
//		}
//	}

}
