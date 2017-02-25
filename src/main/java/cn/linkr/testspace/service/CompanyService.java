package cn.linkr.testspace.service;

import cn.linkr.testspace.entity.SysCompany;
import cn.linkr.testspace.vo.CompanyVo;
import cn.linkr.testspace.vo.DocumentVo;



public interface CompanyService extends BaseService {

	CompanyVo genVo(SysCompany po);

	SysCompany save(CompanyVo vo);



}
