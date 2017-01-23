package cn.linkr.events.service.impl;

import org.springframework.stereotype.Service;

import cn.linkr.events.entity.EvtGuest;
import cn.linkr.events.entity.SysCompany;
import cn.linkr.events.service.CompanyService;
import cn.linkr.events.util.BeanUtilEx;
import cn.linkr.events.vo.CompanyVo;
import cn.linkr.events.vo.GuestVo;

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
