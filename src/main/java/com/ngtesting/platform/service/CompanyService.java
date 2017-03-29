package com.ngtesting.platform.service;

import com.ngtesting.platform.entity.SysCompany;
import com.ngtesting.platform.vo.CompanyVo;
import com.ngtesting.platform.vo.DocumentVo;



public interface CompanyService extends BaseService {

	CompanyVo genVo(SysCompany po);

	SysCompany save(CompanyVo vo);



}
