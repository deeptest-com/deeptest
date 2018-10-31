package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuField;
import com.ngtesting.platform.model.IsuFieldDefine;

import java.util.List;

public interface IssueFieldDao {
    List<IsuFieldDefine> listDefaultFilter();
    List<IsuFieldDefine> listDefaultField();

    List<IsuField> listOrgField(Integer orgId);
    List<IsuField> listOrgFieldDetail(Integer orgId);
}
