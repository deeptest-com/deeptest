package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuField;
import com.ngtesting.platform.model.IsuFieldDefine;
import org.apache.ibatis.annotations.Param;

import java.util.List;
import java.util.Map;

public interface IssueFieldDao {
    List<IsuFieldDefine> listDefaultFilter();
    List<IsuFieldDefine> listDefaultField();

    IsuField getSysField(@Param("id") Integer id, @Param("orgId") Integer orgId);
    IsuField getCustField(@Param("id") Integer id, @Param("orgId") Integer orgId);

    List<Map> fetchInputMap();

    List<Map> listInput();
    List<Map> listType();
}
