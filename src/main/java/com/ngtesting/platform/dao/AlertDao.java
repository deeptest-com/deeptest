package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstAlert;
import com.ngtesting.platform.model.TstUser;
import org.apache.ibatis.annotations.Param;

import java.util.Date;
import java.util.List;

public interface AlertDao {
    List<TstAlert> query(@Param("userId") Integer userId, @Param("isRead") Boolean isRead);

    List<TstAlert> scanAlerts(@Param("userId") Integer userId,
                              @Param("startTimeOfToday") Date startTimeOfToday, @Param("endTimeOfToday") Date endTimeOfToday);

    void create(TstAlert po);
    void update(TstAlert po);

    void markAllRead(@Param("ids") String ids, @Param("userId") Integer userId);

    void removeOldIfNeeded(@Param("taskId") Integer taskId, @Param("assignees") List<TstUser> assignees);
}
