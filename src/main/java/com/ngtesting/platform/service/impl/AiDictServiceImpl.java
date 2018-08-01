package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.model.AiDict;
import com.ngtesting.platform.service.AiDictService;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.stereotype.Service;

import java.io.File;
import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

@Service
public class AiDictServiceImpl extends BaseServiceImpl implements AiDictService {
	private static final Log log = LogFactory.getLog(AiDictServiceImpl.class);

	@Override
    public void dictPers() {
//        // String dirPath = System.getProperties().getProperty("user.home") + "/work/dict/import";
//        String dirPath = "/home/aaron/work/dict/import";
//
//		File dictParentDir = new File(dirPath);
//		if(dictParentDir.isDirectory()){
//			File[] fileArray = dictParentDir.listFiles();
//			for(File dictFile : fileArray ){
//				String catrgory = dirPath;
//				FileReader reader = null;
//	            try {
//	                reader = new FileReader(dictFile);
//	                BufferedReader br = new BufferedReader(reader);
//	                String str = null;
//	                int i = 0;
//	                while ((str = br.readLine()) != null) {
//	                	if (i++ > 1000) {
//	                		break;
//	                	}
//	                    String[] arr;
//						if (str.indexOf(",") > -1) {
//							arr = str.trim().split(",");
//						} else if (str.indexOf(":") > -1) {
//							arr = str.trim().split(":");
//						} else {arr = str.trim().split(":");
//							arr = str.trim().split(":");
//						}
//
//	                    if (arr.length == 0) {
//	                        continue;
//	                    }
//
//	                    AiDict po = find(catrgory, arr[0]);
//	                    if (po != null) {
//	                    	continue;
//	                    }
//
//	                    String synonym = null;
//	                    if (arr.length > 1) {
//	                    	synonym = arr[1];
//	                    }
//
//	                    String skillId = null;
//	                    String[] arr2 = dictFile.getName().split("-");
//	                    if (arr2.length > 1) {
//	                    	skillId = arr2[0];
//	                    	catrgory = arr2[1].replace(".txt", "").replace(".csv", "");
//	                    }
//						AiDict dict = new AiDict(skillId, catrgory, arr[0], synonym);
//	                    saveOrUpdate(dict);
//	                }
//
//	                br.close();
//	                reader.close();
//
//	            } catch (Exception e) {
//	                e.printStackTrace();
//	            }
//			}
//		}
    }

    public List traverseFolder(String path, List<String> fileList) {
        File file = new File(path);
        if (file.exists()) {
            File[] files = file.listFiles();
            if (files.length == 0) {
                System.out.println("文件夹是空的!");
            } else {
                for (File file2 : files) {
                    if (file2.isDirectory()) {
                        System.out.println("文件夹:" + file2.getAbsolutePath());
                        traverseFolder(file2.getAbsolutePath(), fileList);
                    } else {
                        System.out.println("文件:" + file2.getAbsolutePath());
                        if (file2.getAbsolutePath().lastIndexOf(".txt") > 0) {
                            fileList.add(file2.getAbsolutePath());
                        }
                    }
                }
            }
        } else {
            System.out.println("文件不存在!");
        }

        return fileList;
    }

	@Override
	public Map<String, List<String>> get(String json) {
		Map<String, List<String>> ret = new HashMap<>();
//
//		Map<String, Integer> map = JSON.parseObject(json, Map.class);
//
//		for (String key : map.keySet()) {
//			Integer numb = map.get(key);
//			key = key.replace("{", "").replace("}", "");
//			String skillId = null;
//            String[] arr2 = key.split("-");
//            if (arr2.length > 1) {
//            	skillId = arr2[0];
//            	key = arr2[1];
//            }
//
//			String sql = "select count(d.id) from ai_dict d where d.category = ?";
//			String keyNew = "sys.城市".equals(key) ?"sys.国内城市": key;
//			Integer count = Integer.valueOf(getBySQL(sql, new String[]{keyNew}).toString());
//			if (numb > count) {
//				numb = count;
//			}
//			if (numb > 1000) {
//				numb = 1000;
//			}
//			int startIndex = count == numb? 0: new Random().nextInt(count - numb);
//
//			DetachedCriteria dc = DetachedCriteria.forClass(AiDict.class);
//
//			if (skillId != null) {
//				dc.add(Restrictions.eq("skillId", skillId));
//			}
//	        dc.add(Restrictions.eq("category", keyNew));
//
//	        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//	        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//	        dc.addOrder(Order.asc("id"));
//	        Page listByPage = findPage(dc, startIndex, numb);
//
//	        List<String> ls = new LinkedList<String>();
//	        for (Object obj: listByPage.getItems()) {
//				AiDict dict = (AiDict) obj;
//	        	ls.add(dict.getPhrase());
//	        }
//	        if (listByPage.getItems().size() == 0) {
//	        	log.info("***词库 '" + key + "' 不在数据库中");
//
//	        	ls = genRegDict(key);
//	        }
//	        ret.put(key, ls);
//		}

		return ret;
	}

	@Override
    public AiDict find(String category, String phrase) {
//        DetachedCriteria dc = DetachedCriteria.forClass(AiDict.class);
//        dc.add(Restrictions.eq("category", category));
//        dc.add(Restrictions.eq("phrase", phrase));
//
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//        dc.addOrder(Order.asc("id"));
//        List<AiDict> pos = findAllByCriteria(dc);
//
//        if (pos.size() > 0) {
//        	return pos.get(0);
//        }

        return null;
	}

	@Override
    public List<String> genRegDict(String dict) {
		List ret = new LinkedList<String>();
        if ("sys.日期".equals(dict)) {
        	ret = new LinkedList<String>() {{
                add("昨天"); add("今天"); add("明天"); add("后天");
                add("三月"); add("上周五"); add("下周一");

                add("二零一六年十二月八号"); add("二零一七年七月二十八号"); add("一七年七月二十八号"); add("七月二十八号");

//                add(DateUtils.FormatDate(new Date(), "yyyyMMdd"));
//                add(DateUtils.FormatDate(new Date(), "MMdd"));
            }};
        } else if ("sys.时间".equals(dict)) {
        	ret = new LinkedList<String>() {{
        		add("上午九点"); add("中午十二点"); add("晚上八点");
        		add("十二点一刻"); add("十九点半"); add("十九点三十分");
            }};
        } else if ("sys.整数".equals(dict)) {
        	ret = new LinkedList<String>() {{
//        		add("一百"); add("二千三百四拾五"); add("幺仟六百七十");
        		add("九"); add("八六七"); add("幺"); add("五九五九");
            }};
        } else if ("sys.股票代码".equals(dict)) {
        	ret = new LinkedList<String>() {{
        		add("SH123456"); add("SZ098765");
        		add("123456"); add("098765");
            }};
        } else if ("sys.相对时间".equals(dict)) {
        	ret = new LinkedList<String>() {{
        		add("一个小时"); add("一个半小时"); add("一小时三十分");
            }};
        } else if ("sys.序列号".equals(dict)) {
        	ret = new LinkedList<String>() {{
        		add("第一个"); add("第一首");
            }};
        } else if ("sys.故事适用年龄".equals(dict)) {
        	ret = new LinkedList<String>() {{
        		add("三岁半"); add("三个月");
            }};
        } else if ("sys.数量".equals(dict)) {
        	ret = new LinkedList<String>() {{
        		add("一个"); add("几个"); add("几首"); add("几集");
        		add("全部"); add("所有");
            }};
        } else if ("sys.年份".equals(dict)) {
        	ret = new LinkedList<String>() {{
        		add("二零一六年"); add("二零一七年"); add("二零一八年");
            }};
        } else if ("sys.阴历日期".equals(dict)) {
        	ret = new LinkedList<String>() {{
        		add("农历三月初一"); add("农历八月初一"); add("农历十二月初一");
            }};
        }

        return ret;
	}

}
