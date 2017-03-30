package com.ngtesting.platform.service;

import com.ngtesting.platform.entity.SysCompany;
import com.ngtesting.platform.vo.CompanyVo;



public interface CompanyService extends BaseService {

	CompanyVo genVo(SysCompany po);

	SysCompany save(CompanyVo vo);



}
