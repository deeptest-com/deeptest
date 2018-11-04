package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuPage;
import com.ngtesting.platform.model.IsuPageElement;
import com.ngtesting.platform.model.IsuPageTab;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface IssuePageDao {

    List<IsuPage> list(Integer orgId);

    IsuPage get(@Param("id") Integer id, @Param("orgId") Integer orgId);

    void save(IsuPage vo);

    Integer update(IsuPage vo);

    Integer delete(@Param("id") Integer id, @Param("orgId") Integer orgId);

    void saveDefaultTab(@Param("pageId") Integer pageId, @Param("orgId") Integer orgId);

    void addTab(IsuPageTab tab);

    void addField(IsuPageElement element);

    IsuPageTab getTab(@Param("tabId") Integer tabId, @Param("orgId") Integer orgId);

    Integer getMaxFieldOrdr(Integer tabId);
}