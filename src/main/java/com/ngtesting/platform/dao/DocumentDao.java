package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.Document;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface DocumentDao {
    List<Document> query(@Param("userId") Integer userId);
}
