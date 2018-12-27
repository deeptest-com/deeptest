package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuField;
import org.apache.ibatis.annotations.Param;

import java.util.List;
import java.util.Map;

public interface IssueDynamicFormDao {
    List<IsuField> listNotUsedField(@Param("orgId") Integer orgId,
                                    @Param("projectId") Integer projectId,
                                    @Param("pageId") Integer pageId,
                                    @Param("sort") String sort);

    List<Map> fetchOrgField(@Param("orgId") Integer orgId,
                            @Param("projectId") Integer projectId,
                            @Param("sort") String sort);
}
