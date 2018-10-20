package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuFieldDefine;

import java.util.List;

public interface IsuFieldDefineDao {
    List<IsuFieldDefine> listFilters();

    List<IsuFieldDefine> listColumns();
}
