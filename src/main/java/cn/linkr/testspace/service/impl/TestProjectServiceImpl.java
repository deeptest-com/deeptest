package cn.linkr.testspace.service.impl;

import java.util.Date;
import java.util.HashMap;
import java.util.HashSet;
import java.util.Iterator;
import java.util.LinkedHashSet;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

import net.sf.ehcache.CacheManager;
import net.sf.ehcache.Element;
import net.sf.ehcache.management.Cache;

import org.apache.commons.lang.StringUtils;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.hibernate.FetchMode;
import org.hibernate.Filter;
import org.hibernate.criterion.CriteriaSpecification;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.cache.annotation.Cacheable;
import org.springframework.stereotype.Service;

import com.alibaba.fastjson.JSONObject;

import cn.linkr.testspace.action.ProjectAction;
import cn.linkr.testspace.dao.BaseDao;
import cn.linkr.testspace.dao.ProjectDao;
import cn.linkr.testspace.entity.EvtEvent;
import cn.linkr.testspace.entity.EvtGuest;
import cn.linkr.testspace.entity.EvtScheduleItem;
import cn.linkr.testspace.entity.EvtSession;
import cn.linkr.testspace.entity.SysCompany;
import cn.linkr.testspace.entity.TestCase;
import cn.linkr.testspace.entity.TestProject;
import cn.linkr.testspace.entity.EvtEvent.EventStatus;
import cn.linkr.testspace.entity.TestProject.ProjectType;
import cn.linkr.testspace.service.GuestService;
import cn.linkr.testspace.service.TestCaseService;
import cn.linkr.testspace.service.TestProjectService;
import cn.linkr.testspace.util.BeanUtilEx;
import cn.linkr.testspace.util.Constant;
import cn.linkr.testspace.util.Constant.TreeNodeType;
import cn.linkr.testspace.util.StringUtil;
import cn.linkr.testspace.vo.GuestVo;
import cn.linkr.testspace.vo.Page;
import cn.linkr.testspace.vo.SessionVo;
import cn.linkr.testspace.vo.TestCaseTreeVo;
import cn.linkr.testspace.vo.TestCaseVo;
import cn.linkr.testspace.vo.TestProjectVo;
import cn.linkr.testspace.vo.UserVo;

@Service
public class TestProjectServiceImpl extends BaseServiceImpl implements
		TestProjectService {

	private static final Log log = LogFactory
			.getLog(TestProjectServiceImpl.class);

	@Autowired
	private ProjectDao projectDao;

	@Override
	// @Cacheable(value="companyProjects",key="#companyId.toString().concat('_').concat(#disabled)")
	public List<TestProjectVo> list(Long companyId, String keywords, String disabled) {
		// CacheManager manager = CacheManager.create();
		// net.sf.ehcache.Cache cache = manager.getCache("companyProjects");
		// String key = companyId + "_" + disabled;
		// Element el = null;
		// if(cache.isKeyInCache(key)){
		// el = cache.get(key);
		// return (Map<String, Object>)el.getObjectValue();
		// }

		DetachedCriteria dc = DetachedCriteria.forClass(TestProject.class);
		
		if (StringUtil.isNotEmpty(disabled)) {
			dc.add(Restrictions.eq("disabled", Boolean.valueOf(disabled)));
		}

		dc.add(Restrictions.eq("type", ProjectType.group));
		dc.add(Restrictions.eq("deleted", Boolean.FALSE));
		dc.addOrder(Order.asc("id"));
		
		dc.setFetchMode("children", FetchMode.JOIN);
		dc.setResultTransformer(CriteriaSpecification.DISTINCT_ROOT_ENTITY); 
		
		Filter filter = getDao().getSession().enableFilter("filter_disabled");
		filter.setParameter("isDeleted", Boolean.valueOf(false));
		List<TestProject> pos = findAllByCriteria(dc);
		getDao().getSession().disableFilter("filter_disabled");

		List<TestProjectVo> vos = this.genVos(pos, keywords, disabled);

		// el = new Element(key, ret);
		// cache.put(el);
		return vos;
	}
	
	@Override
	public List<TestProjectVo> listGroups(Long companyId) {
		DetachedCriteria dc = DetachedCriteria.forClass(TestProject.class);

		dc.add(Restrictions.eq("type", ProjectType.group));
		dc.add(Restrictions.eq("disabled", false));
		
		dc.add(Restrictions.eq("deleted", Boolean.FALSE));
		dc.addOrder(Order.asc("id"));
		
		List<TestProject> pos = findAllByCriteria(dc);

		List<TestProjectVo> vos = this.genGroupVos(pos);

		return vos;
	}

	@Override
	public TestProject getDetail(Long id) {
		if (id == null) {
			return null;
		}
		TestProject po = (TestProject) get(TestProject.class, id);

		return po;
	}

	@Override
	public TestProject save(TestProjectVo vo, UserVo user) {
		if (vo == null) {
			return null;
		}

		boolean isNew = vo.getId() == null;
		TestProject po = new TestProject();
		if (!isNew) {
			po = (TestProject) get(TestProject.class, vo.getId());
		} else {
			po.setCompanyId(user.getCompanyId());
		}
		po.setParentId(vo.getParentId());
		po.setName(vo.getName());
		po.setDescr(vo.getDescr());
		po.setType(ProjectType.getEnum(vo.getType()));
		po.setDisabled(vo.getDisabled());

		saveOrUpdate(po);
		
		// 项目被启用，启用父项目
		TestProject parent = po.getParent();
		if (parent != null && po.getType().equals(ProjectType.project) && !po.getDisabled() && parent.getDisabled()) { 
			parent.setDisabled(false);
			saveOrUpdate(parent);
		}
		
		// 项目组被禁用，禁用子项目
		if (po.getType().equals(ProjectType.group) && po.getDisabled()) {
			for (TestProject child : po.getChildren()) {
				child.setDisabled(true);
				saveOrUpdate(child);
			}
			
		}

		// this.removeCache(user.getCompanyId());

		return po;
	}

	@Override
	public Boolean delete(Long id, Long userId) {
		if (id == null) {
			return null;
		}

		TestProject po = (TestProject) get(TestProject.class, id);
		po.setDeleted(true);
		saveOrUpdate(po);
		
		// 项目组被删除，删除子项目
		if (po.getType().equals(ProjectType.group)) {
			for (TestProject child : po.getChildren()) {
				child.setDeleted(true);
				saveOrUpdate(child);
			}
			
		}

		return true;
	}

	// @Override
	// public void removeCache(Long companyId) {
	// CacheManager manager = CacheManager.create();
	// net.sf.ehcache.Cache cache = manager.getCache("companyProjects");
	// String prefix = companyId + "_";
	// if(cache.isKeyInCache(prefix + "true")){
	// cache.remove(prefix + "true");
	// }
	// if(cache.isKeyInCache(prefix)){
	// cache.remove(prefix);
	// }
	// }

	@Override
	public List<TestProjectVo> genVos(List<TestProject> pos, String keywords, String disabled) {
		List<TestProjectVo> voList = new LinkedList<TestProjectVo>();
		for (TestProject po : pos) {
			TestProjectVo vo = genVo(po);
			voList.add(vo);
			
			List<TestProjectVo> voList2 = new LinkedList<TestProjectVo>();
			List<TestProject> children = po.getChildren();
			for (TestProject child : children) {
				if ( (StringUtil.IsEmpty(keywords) || child.getName().toLowerCase().indexOf(keywords.toLowerCase()) > -1) 
						&& ( StringUtil.IsEmpty(disabled) || child.getDisabled() == Boolean.valueOf(disabled)) ) {
					TestProjectVo childVo = genVo(child);
					voList2.add(childVo);
				}
			}
			vo.setChildrenNumb(voList2.size());
			voList.addAll(voList2);
		}
		
		return voList;
	}
	@Override
	public List<TestProjectVo> genGroupVos(List<TestProject> pos) {
		List<TestProjectVo> voList = new LinkedList<TestProjectVo>();
		for (TestProject po : pos) {
				TestProjectVo vo = genVo(po);
				voList.add(vo);
		}
		
		return voList;
	}

	@Override
	public TestProjectVo genVo(TestProject po) {
		if (po == null) {
			return null;
		}
		TestProjectVo vo = new TestProjectVo();
		BeanUtilEx.copyProperties(vo, po);
		if (po.getParentId() == null) {
			vo.setParentId(null);
		}
		return vo;
	}

}
