package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuField;
import org.apache.ibatis.annotations.Param;

import java.util.List;
import java.util.Map;

public interface IssueDynamicFormDao {
    List<IsuField> listTabNotUsedField(@Param("orgId") Integer orgId,
                                       @Param("projectId") Integer projectId,
                                       @Param("tabId") Integer tabId);

    List<Map> fetchOrgField(@Param("orgId") Integer orgId,
                            @Param("projectId") Integer projectId);
}
