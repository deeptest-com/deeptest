package cn.linkr.events.service;

import cn.linkr.events.entity.SysCompany;
import cn.linkr.events.vo.CompanyVo;
import cn.linkr.events.vo.DocumentVo;



public interface CompanyService extends BaseService {

	CompanyVo genVo(SysCompany po);

	SysCompany save(CompanyVo vo);



}
