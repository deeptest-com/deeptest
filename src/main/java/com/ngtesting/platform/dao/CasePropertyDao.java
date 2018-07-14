package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstMsg;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface CasePropertyDao {
    List<TstMsg> query(@Param("userId") Integer userId);
}
