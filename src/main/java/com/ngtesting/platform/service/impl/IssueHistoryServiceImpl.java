package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.model.TstHistory;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.IssueHistoryService;
import org.springframework.stereotype.Service;

import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;

@Service
public class IssueHistoryServiceImpl extends BaseServiceImpl implements IssueHistoryService {

	@Override
	public List<TstHistory> list(Integer projectId, String projectType) {
//		DetachedCriteria dc = DetachedCriteria.forClass(TstHistory.class);
//
//		dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//		dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//
//		if (projectType.equals(TestProject.ProjectType.project.toString())) {
//			dc.add(Restrictions.eq("projectId", projectId));
//		} else {
//			dc.createAlias("project", "project");
//			dc.add(Restrictions.eq("project.parentId", projectId));
//		}
//
//		dc.addOrder(Order.desc("createTime"));
//
//		List<TstHistory> ls = findAllByCriteria(dc);

		return null;
	}

    @Override
    public TstHistory getById(Integer id) {
//		TstHistory po = (TstHistory) getDetail(TstHistory.class, id);
//		TstHistory vo = genVo(po);
//
//        return vo;

		return null;
    }

    @Override
    public TstHistory create(Integer projectId, TstUser optUser, String action,
							  TstHistory.TargetType entityType, Integer entityId, String name) {
        TstHistory history = new TstHistory();

//        history.setTitle("用户" + StringUtil.highlightDict(optUser.getName())
//                + action + entityType.name + StringUtil.highlightDict(name));
//        history.setProjectId(projectId);
//        history.setEntityId(entityId);
//        history.setEntityType(entityType);
//        history.setUserId(optUser.getCode());
//        saveOrUpdate(history);

        return history;
    }

	@Override
	public Map<String, List<TstHistory>> genVosByDate(List<TstHistory> historyPos) {
		Map<String, List<TstHistory>> map = new LinkedHashMap();
//		for(TstHistory his: historyPos) {
//            Date createDate = his.getCreateTime();
//            String date = DateUtils.FormatDate(createDate, "yyyy-MM-dd");
//            if (!map.containsKey(date)) {
//                map.put(date, new LinkedList());
//            }
//            map.getDetail(date).add(genVo(his));
//		}
		return map;
	}

}

