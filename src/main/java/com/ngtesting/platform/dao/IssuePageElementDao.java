package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.CustomField;
import com.ngtesting.platform.model.IsuPageElement;
import org.apache.ibatis.annotations.Param;

import java.util.List;
import java.util.Map;

public interface IssuePageElementDao {
    void save(IsuPageElement elm);

    void updateProp(@Param("id") Integer id,
                    @Param("prop") String prop,
                    @Param("val") String val,
                    @Param("orgId") Integer orgId);

    void updateFromCustomField(CustomField vo);

    void saveOrdrs(@Param("maps") List<Map> maps,
                   @Param("pageId") Integer pageId,
                   @Param("orgId") Integer orgId);

    void removeOthers(@Param("maps") List<Map> maps,
                      @Param("pageId") Integer pageId,
                      @Param("orgId") Integer orgId);

    List<IsuPageElement> listElementByPageId(@Param("id") Integer id);
}