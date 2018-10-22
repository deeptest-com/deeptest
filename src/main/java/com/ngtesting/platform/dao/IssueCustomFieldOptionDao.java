package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuCustomFieldOption;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface IssueCustomFieldOptionDao {
    List<IsuCustomFieldOption> listByFieldId(@Param("fieldId") Integer fieldId);
    IsuCustomFieldOption get(@Param("id") Integer id);

    void save(IsuCustomFieldOption vo);
    void saveAll(@Param("fieldId") Integer fieldId, @Param("list") List<IsuCustomFieldOption> options);

    void update(IsuCustomFieldOption vo);
    void delete(@Param("id") Integer id);

    void setOrder(@Param("id") Integer id,
                  @Param("ordr") Integer ordr);

    IsuCustomFieldOption getPrev(@Param("ordr") Integer ordr,
                                 @Param("fieldId") Integer fieldId);

    IsuCustomFieldOption getNext(@Param("ordr") Integer ordr,
                                 @Param("fieldId") Integer fieldId);

    Integer getMaxOrder(@Param("fieldId") Integer fieldId);

}
