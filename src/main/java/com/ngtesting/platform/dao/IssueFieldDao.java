package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuField;
import com.ngtesting.platform.model.IsuFieldDefine;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface IssueFieldDao {
    List<IsuFieldDefine> listDefaultFieldInColumns();

    IsuField getSysField(@Param("id") Integer id, @Param("orgId") Integer orgId);
    IsuField getCustField(@Param("id") Integer id, @Param("orgId") Integer orgId);

}
