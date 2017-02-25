package cn.linkr.testspace.service.impl;

import org.springframework.stereotype.Service;

import cn.linkr.testspace.entity.EvtGuest;
import cn.linkr.testspace.entity.SysCompany;
import cn.linkr.testspace.service.CompanyService;
import cn.linkr.testspace.util.BeanUtilEx;
import cn.linkr.testspace.vo.CompanyVo;
import cn.linkr.testspace.vo.GuestVo;

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
