package com.ngtesting.platform.service.impl;

import org.springframework.stereotype.Service;

import com.ngtesting.platform.entity.EvtGuest;
import com.ngtesting.platform.entity.SysCompany;
import com.ngtesting.platform.service.CompanyService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.vo.CompanyVo;
import com.ngtesting.platform.vo.GuestVo;

@Service
public class CompanyServiceImpl extends BaseServiceImpl implements CompanyService {

	@Override
	public CompanyVo genVo(SysCompany po) {
		CompanyVo vo = new CompanyVo();
		BeanUtilEx.copyProperties(vo, po);
		return vo;
	}

	@Override
	public SysCompany save(CompanyVo vo) {
		if (vo == null) {
			return null;
		}
		
		SysCompany po = new SysCompany();
		if (vo.getId() != null) {
			po = (SysCompany) get(SysCompany.class, vo.getId());
		}
		
		po.setName(vo.getName());
		
		saveOrUpdate(po);
		return po;
	}

    
}
