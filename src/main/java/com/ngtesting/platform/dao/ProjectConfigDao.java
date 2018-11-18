package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstProject;
import org.apache.ibatis.annotations.Param;

public interface ProjectConfigDao {
    TstProject get(@Param("id") Integer id);
}
