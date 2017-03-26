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
	
	private static final Log log = LogFactory.getLog(TestProjectServiceImpl.class);

	@Autowired
	private ProjectDao projectDao;
	
	@Override
//	@Cacheable(value="companyProjects",key="#companyId.toString().concat('_').concat(#isActive)")
	public Map<String, Object> listCache(Long companyId, String isActive) {
		CacheManager manager = CacheManager.create();
		net.sf.ehcache.Cache cache = manager.getCache("companyProjects");
		String key = companyId + "_" + isActive;
		Element el = null;
//        if(cache.isKeyInCache(key)){
//        	el = cache.get(key);
//            return (Map<String, Object>)el.getObjectValue();
//        }
        
		DetachedCriteria dc = DetachedCriteria.forClass(TestProject.class);

		if (StringUtil.isNotEmpty(isActive)) {
			dc.add(Restrictions.eq("isActive", Boolean.valueOf(isActive)));
		}

		dc.add(Restrictions.eq("deleted", Boolean.FALSE));
		dc.add(Restrictions.eq("disabled", Boolean.FALSE));
		dc.addOrder(Order.asc("path"));
		dc.addOrder(Order.asc("id"));
		List<TestProject> pos = findAllByCriteria(dc);
		
		Map<String, Integer> out = new HashMap<String, Integer>();
		LinkedList<TestProjectVo> vos = this.genVos(pos, out);

		Map<String, Object> ret = new HashMap<String, Object>();
		ret.put("models", vos);
		ret.put("maxLevel", out.get("maxLevel"));
		
		el = new Element(key, ret);
        cache.put(el);
		return ret;
	}

	@Override
    public TestProject getDetail(Long id) {
    	TestProject po = (TestProject) get(TestProject.class, id);
		
		return po;
    }
	
	@Override
	public TestProject save(TestProjectVo vo, UserVo user) {
		if (vo == null) {
			return null;
		}
		
		TestProject po = new TestProject();
		if (vo.getId() != null) {
			po = (TestProject) get(TestProject.class, vo.getId());
		}
		
		po.setName(vo.getName());
		po.setDescr(vo.getDescr());
		
		saveOrUpdate(po);
		
		if (vo.getParentId() != po.getParentId()) {
			getDao().moveNode("tst_project", po.getId(), vo.getParentId());
		}
		
		if (vo.getIsActive() != po.getIsActive()) {
			getDao().updateNode("tst_project", po.getId(), "is_active", vo.getIsActive().toString());
		}
		
		this.removeCache(user.getCompanyId());
		
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
		
		getDao().updateNode("tst_project", po.getId(), "deleted", "true");
		
		return true;
	}
	
	@Override
	public void removeCache(Long companyId) {
		CacheManager manager = CacheManager.create();
        net.sf.ehcache.Cache cache = manager.getCache("companyProjects");
        String prefix = companyId + "_";
        if(cache.isKeyInCache(prefix + "true")){
            cache.remove(prefix + "true");
        }
        if(cache.isKeyInCache(prefix)){
            cache.remove(prefix);
        }
	}
	
	@Override
	public LinkedList<TestProjectVo> genVos(List<TestProject> pos, Map<String, Integer> ret) {
		TestProjectVo root = null;
		int maxLevel = 0;
		
		Map<Long, TestProjectVo> nodeMap = new HashMap<Long, TestProjectVo>();
	      for (TestProject po : pos) {
	        	Long id = po.getId();
	        	TreeNodeType type = po.getType();
	        	Long pid = po.getParentId();
	        	
	        	TestProjectVo newNode = genVo(po);
	        	
	        	nodeMap.put(id, newNode);
	        	
	        	if (type.equals(Constant.TreeNodeType.root)) {
	        		root = newNode;
	        		continue;
	        	}
	        	
	        	TestProjectVo pvo = nodeMap.get(pid);
	        	LinkedList<TestProjectVo> children = pvo.getChildren();
	        	children.add(newNode);
	        	
	        	if (po.getLevel() > maxLevel) {
					maxLevel = po.getLevel();
				}
	        }
	        
	      LinkedList<TestProjectVo> voList = new LinkedList<TestProjectVo>();
	      this.toOrderList(root, voList);
	      this.removeChildren(voList);
		
        ret.put("maxLevel", maxLevel);
        return voList;
	}

	@Override
	public TestProjectVo genVo(TestProject po) {
		TestProjectVo vo = new TestProjectVo();
		BeanUtilEx.copyProperties(vo, po);
		if (po.getParentId() == null) {
			vo.setParentId(null);
		}
		return vo;
	}

	@Override
	public void toOrderList(TestProjectVo root, LinkedList<TestProjectVo> resultList) {
		resultList.add(root);
		
		Iterator<TestProjectVo> it = root.getChildren().iterator();
		while (it.hasNext()) {
			TestProjectVo vo = it.next();
        	this.toOrderList(vo, resultList);
		}
	}
	
	@Override
	public void removeChildren(LinkedList<TestProjectVo> resultList) {
		Iterator<TestProjectVo> it = resultList.iterator();
		while (it.hasNext()) {
			TestProjectVo vo = it.next();
			vo.setChildren(null);
		}
	}

	@Override
	public LinkedList<TestProjectVo> removeChildren(LinkedList<TestProjectVo> linkedList, TestProjectVo curr) {
		LinkedList<TestProjectVo> ret = new LinkedList<TestProjectVo>();
		
		Iterator<TestProjectVo> it = linkedList.iterator();
		
		while (it.hasNext()) {
			TestProjectVo vo = it.next();
			if (curr.getId() != vo.getId() && vo.getPath().indexOf("/" + curr.getId() + "/") == -1) {
				ret.add(vo);
			}
		}
		
		return ret;
	}
}
