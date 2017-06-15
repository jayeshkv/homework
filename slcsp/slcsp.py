#!/usr/bin/python2.7
# coding: utf-8

# In[419]:

import numpy as np
import pandas as pd

slcsp_csv = pd.read_csv('slcsp.csv')
plans_csv = pd.read_csv('plans.csv')
zips_csv = pd.read_csv('zips.csv')

plans_csv['rate_area'] = list(zip(plans_csv.state, plans_csv.rate_area))
plans_csv= plans_csv.drop('state',1)
plans_csv = plans_csv[plans_csv.metal_level=='Silver']

zips_csv['rate_area'] = list(zip(zips_csv.state, zips_csv.rate_area)) #joining state & rate_area
zips_csv= zips_csv.drop('state',1) #dropping state from DataFrame

mergedData =  pd.merge(plans_csv, zips_csv, on=['rate_area']) #merging based on rate_area (GA, 7)

slcsp_csv= slcsp_csv.drop('rate',1)#dropping rate from DF
zip_area_rate = pd.merge(slcsp_csv,mergedData[['rate_area','rate', 'zipcode']], on='zipcode', how='inner')#merging with slcsp_csv based on zipcode #dropping duplicates based on both the cols

#setting up flags for 2 rate areas
handle_R_A = zip_area_rate.groupby('zipcode').rate_area.transform('nunique') > 1
zip_area_rate.loc[handle_R_A, 'test_flag']='T'
zip_area_rate = zip_area_rate[zip_area_rate.test_flag != 'T'].reset_index()
zip_area_rate= zip_area_rate.drop_duplicates(["zipcode","rate"]) #dropping duplicates based on both the cols

fin = zip_area_rate.groupby(['zipcode'])['rate'].nsmallest(2).groupby(level='zipcode').last().reset_index()
finalCsv = pd.merge(slcsp_csv, fin, on='zipcode', how='outer')#finally merging with slcsp_csv for filtering 
finalCsv=  finalCsv.fillna('')

finalCsv = pd.DataFrame(finalCsv, columns=['zipcode','rate'])
fileName = 'modifiedCsv.csv'
finalCsv.to_csv(fileName, index=False, encoding='utf-8')

