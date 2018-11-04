package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuField;
import com.ngtesting.platform.model.IsuFieldDefine;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface IssueFieldDao {
    List<IsuFieldDefine> listDefaultFilter();
    List<IsuFieldDefine> listDefaultField();

    List<IsuField> listOrgField(@Param("orgId") Integer orgId, @Param("tabId") Integer tabId);
}
