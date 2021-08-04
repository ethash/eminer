/*
 * Copyright 2011-2012 Con Kolivas
 *
 * This program is free software; you can redistribute it and/or modify it
 * under the terms of the GNU General Public License as published by the Free
 * Software Foundation; either version 3 of the License, or (at your option)
 * any later version.  See COPYING for more details.
 */
#include "adlbridge.h"

#if (defined(__linux__) || defined (WIN32))

bool adl_active = false;

// Memory allocation function
void * __stdcall ADL_Main_Memory_Alloc(int iSize)
{
  void *lpBuffer = malloc(iSize);

  return lpBuffer;
}

// Optional Memory de-allocation function
void __stdcall ADL_Main_Memory_Free (void **lpBuffer)
{
  if (*lpBuffer != NULL) {
    free (*lpBuffer);
    *lpBuffer = NULL;
  }
}

#if defined (__linux__)
// equivalent functions in linux
void *GetProcAddress(void *pLibrary, const char *name)
{
  return dlsym( pLibrary, name);
}
#endif

 ADL_MAIN_CONTROL_CREATE     ADL_Main_Control_Create;
 ADL_MAIN_CONTROL_DESTROY    ADL_Main_Control_Destroy;
 ADL_ADAPTER_NUMBEROFADAPTERS_GET  ADL_Adapter_NumberOfAdapters_Get;
 ADL_ADAPTER_ADAPTERINFO_GET   ADL_Adapter_AdapterInfo_Get;
 ADL_ADAPTER_ID_GET      ADL_Adapter_ID_Get;
 ADL_MAIN_CONTROL_REFRESH    ADL_Main_Control_Refresh;
 ADL_ADAPTER_VIDEOBIOSINFO_GET   ADL_Adapter_VideoBiosInfo_Get;
 ADL_DISPLAY_DISPLAYINFO_GET   ADL_Display_DisplayInfo_Get;
 ADL_ADAPTER_ACCESSIBILITY_GET   ADL_Adapter_Accessibility_Get;

 ADL_OVERDRIVE_CAPS      ADL_Overdrive_Caps;

 ADL_OVERDRIVE5_TEMPERATURE_GET    ADL_Overdrive5_Temperature_Get;
 ADL_OVERDRIVE5_CURRENTACTIVITY_GET  ADL_Overdrive5_CurrentActivity_Get;
 ADL_OVERDRIVE5_ODPARAMETERS_GET   ADL_Overdrive5_ODParameters_Get;
 ADL_OVERDRIVE5_FANSPEEDINFO_GET   ADL_Overdrive5_FanSpeedInfo_Get;
 ADL_OVERDRIVE5_FANSPEED_GET   ADL_Overdrive5_FanSpeed_Get;
 ADL_OVERDRIVE5_FANSPEED_SET   ADL_Overdrive5_FanSpeed_Set;
 ADL_OVERDRIVE5_ODPERFORMANCELEVELS_GET  ADL_Overdrive5_ODPerformanceLevels_Get;
 ADL_OVERDRIVE5_ODPERFORMANCELEVELS_SET  ADL_Overdrive5_ODPerformanceLevels_Set;
 ADL_OVERDRIVE5_POWERCONTROL_GET   ADL_Overdrive5_PowerControl_Get;
 ADL_OVERDRIVE5_POWERCONTROL_SET   ADL_Overdrive5_PowerControl_Set;
 ADL_OVERDRIVE5_FANSPEEDTODEFAULT_SET  ADL_Overdrive5_FanSpeedToDefault_Set;

 ADL_OVERDRIVE6_CAPABILITIES_GET   ADL_Overdrive6_Capabilities_Get;
 ADL_OVERDRIVE6_FANSPEED_GET   ADL_Overdrive6_FanSpeed_Get;
 ADL_OVERDRIVE6_THERMALCONTROLLER_CAPS ADL_Overdrive6_ThermalController_Caps;
 ADL_OVERDRIVE6_TEMPERATURE_GET    ADL_Overdrive6_Temperature_Get;
 ADL_OVERDRIVE6_STATEINFO_GET    ADL_Overdrive6_StateInfo_Get;
 ADL_OVERDRIVE6_CURRENTSTATUS_GET  ADL_Overdrive6_CurrentStatus_Get;
 ADL_OVERDRIVE6_POWERCONTROL_CAPS  ADL_Overdrive6_PowerControl_Caps;
 ADL_OVERDRIVE6_POWERCONTROLINFO_GET ADL_Overdrive6_PowerControlInfo_Get;
 ADL_OVERDRIVE6_POWERCONTROL_GET   ADL_Overdrive6_PowerControl_Get;
 ADL_OVERDRIVE6_FANSPEED_SET   ADL_Overdrive6_FanSpeed_Set;
 ADL_OVERDRIVE6_STATE_SET    ADL_Overdrive6_State_Set;
 ADL_OVERDRIVE6_POWERCONTROL_SET   ADL_Overdrive6_PowerControl_Set;

#if defined (__linux__)
  void *hDLLA;  // Handle to .so library
#else
  HINSTANCE hDLLA;   // Handle to DLL
#endif

int iNumberAdapters;
LPAdapterInfo lpInfo = NULL;


char *adl_error_desc(int error)
{
  char *result;
  switch(error)
  {
    case ADL_ERR:
      result = "Generic error (escape call failed?)";
      break;
    case ADL_ERR_NOT_INIT:
      result = "ADL not initialized";
      break;
    case ADL_ERR_INVALID_PARAM:
      result = "Invalid parameter";
      break;
    case ADL_ERR_INVALID_PARAM_SIZE:
      result = "Invalid parameter size";
      break;
    case ADL_ERR_INVALID_ADL_IDX:
      result = "Invalid ADL index";
      break;
    case ADL_ERR_INVALID_CONTROLLER_IDX:
      result = "Invalid controller index";
      break;
    case ADL_ERR_INVALID_DIPLAY_IDX:
      result = "Invalid display index";
      break;
    case ADL_ERR_NOT_SUPPORTED:
      result = "Function not supported by the driver";
      break;
    case ADL_ERR_NULL_POINTER:
      result = "Null Pointer error";
      break;
    case ADL_ERR_DISABLED_ADAPTER:
      result = "Disabled adapter, can't make call";
      break;
    case ADL_ERR_INVALID_CALLBACK:
      result = "Invalid callback";
      break;
    case ADL_ERR_RESOURCE_CONFLICT:
      result = "Display resource conflict";
      break;
    case ADL_ERR_SET_INCOMPLETE:
      result = "Failed to update some of the values";
      break;
    case ADL_ERR_NO_XDISPLAY:
      result = "No Linux XDisplay in Linux Console environment";
      break;
    default:
      result = "Unhandled error";
      break;
  }
  return result;
}

inline void lock_adl(void)
{
  pthread_mutex_lock(&adl_lock);
}

inline void unlock_adl(void)
{
  pthread_mutex_unlock(&adl_lock);
}

char *gpu_name(int gpu)
{
    struct gpu_adl *ga;
    char *name = "";

    if (!gpus[gpu].has_adl || !adl_active)
      return name;

    return gpus[gpu].name;
}

int gpu_index(int bus_number)
{
    int i;

    for (i = 0; i < MAX_GPUDEVICES; i++)
    {
        if (!gpus[i].has_adl || !adl_active)
            continue;

        if (gpus[i].adl.iBusNumber == bus_number)
            return i;
    }

    return -1;
}

/* This looks for the twin GPU that has the fanspeed control of a non fanspeed
 * control GPU on dual GPU cards */
bool fanspeed_twin(struct gpu_adl *ga, struct gpu_adl *other_ga)
{
  if (!other_ga->has_fanspeed)
    return false;
  if (abs(ga->iBusNumber - other_ga->iBusNumber) != 1)
    return false;
  if (strcmp(ga->strAdapterName, other_ga->strAdapterName))
    return false;
  return true;
}

bool init_overdrive5()
{
  ADL_Overdrive5_Temperature_Get = (ADL_OVERDRIVE5_TEMPERATURE_GET) GetProcAddress(hDLLA,"ADL_Overdrive5_Temperature_Get");
  ADL_Overdrive5_CurrentActivity_Get = (ADL_OVERDRIVE5_CURRENTACTIVITY_GET) GetProcAddress(hDLLA, "ADL_Overdrive5_CurrentActivity_Get");
  ADL_Overdrive5_ODParameters_Get = (ADL_OVERDRIVE5_ODPARAMETERS_GET) GetProcAddress(hDLLA, "ADL_Overdrive5_ODParameters_Get");
  ADL_Overdrive5_FanSpeedInfo_Get = (ADL_OVERDRIVE5_FANSPEEDINFO_GET) GetProcAddress(hDLLA, "ADL_Overdrive5_FanSpeedInfo_Get");
  ADL_Overdrive5_FanSpeed_Get = (ADL_OVERDRIVE5_FANSPEED_GET) GetProcAddress(hDLLA, "ADL_Overdrive5_FanSpeed_Get");
  ADL_Overdrive5_FanSpeed_Set = (ADL_OVERDRIVE5_FANSPEED_SET) GetProcAddress(hDLLA, "ADL_Overdrive5_FanSpeed_Set");
  ADL_Overdrive5_ODPerformanceLevels_Get = (ADL_OVERDRIVE5_ODPERFORMANCELEVELS_GET) GetProcAddress(hDLLA, "ADL_Overdrive5_ODPerformanceLevels_Get");
  ADL_Overdrive5_ODPerformanceLevels_Set = (ADL_OVERDRIVE5_ODPERFORMANCELEVELS_SET) GetProcAddress(hDLLA, "ADL_Overdrive5_ODPerformanceLevels_Set");
  ADL_Overdrive5_PowerControl_Get = (ADL_OVERDRIVE5_POWERCONTROL_GET) GetProcAddress(hDLLA, "ADL_Overdrive5_PowerControl_Get");
  ADL_Overdrive5_PowerControl_Set = (ADL_OVERDRIVE5_POWERCONTROL_SET) GetProcAddress(hDLLA, "ADL_Overdrive5_PowerControl_Set");
  ADL_Overdrive5_FanSpeedToDefault_Set = (ADL_OVERDRIVE5_FANSPEEDTODEFAULT_SET) GetProcAddress(hDLLA, "ADL_Overdrive5_FanSpeedToDefault_Set");

  if (!ADL_Overdrive5_Temperature_Get || !ADL_Overdrive5_CurrentActivity_Get ||
    !ADL_Overdrive5_ODParameters_Get || !ADL_Overdrive5_FanSpeedInfo_Get ||
    !ADL_Overdrive5_FanSpeed_Get || !ADL_Overdrive5_FanSpeed_Set ||
    !ADL_Overdrive5_ODPerformanceLevels_Get || !ADL_Overdrive5_ODPerformanceLevels_Set ||
    !ADL_Overdrive5_PowerControl_Get || !ADL_Overdrive5_PowerControl_Set ||
    !ADL_Overdrive5_FanSpeedToDefault_Set) {
      PRINTFLN("ATI ADL Overdrive5's API is missing or broken.");
      return false;
  } else {
    PRINTFLN("ATI ADL Overdrive5 API found.");
  }

  return true;
}

bool init_overdrive6()
{
  ADL_Overdrive6_FanSpeed_Get = (ADL_OVERDRIVE6_FANSPEED_GET) GetProcAddress(hDLLA,"ADL_Overdrive6_FanSpeed_Get");
  ADL_Overdrive6_ThermalController_Caps = (ADL_OVERDRIVE6_THERMALCONTROLLER_CAPS)GetProcAddress (hDLLA, "ADL_Overdrive6_ThermalController_Caps");
  ADL_Overdrive6_Temperature_Get = (ADL_OVERDRIVE6_TEMPERATURE_GET)GetProcAddress (hDLLA, "ADL_Overdrive6_Temperature_Get");
  ADL_Overdrive6_Capabilities_Get = (ADL_OVERDRIVE6_CAPABILITIES_GET)GetProcAddress(hDLLA, "ADL_Overdrive6_Capabilities_Get");
  ADL_Overdrive6_StateInfo_Get = (ADL_OVERDRIVE6_STATEINFO_GET)GetProcAddress(hDLLA, "ADL_Overdrive6_StateInfo_Get");
  ADL_Overdrive6_CurrentStatus_Get = (ADL_OVERDRIVE6_CURRENTSTATUS_GET)GetProcAddress(hDLLA, "ADL_Overdrive6_CurrentStatus_Get");
  ADL_Overdrive6_PowerControl_Caps = (ADL_OVERDRIVE6_POWERCONTROL_CAPS)GetProcAddress(hDLLA, "ADL_Overdrive6_PowerControl_Caps");
  ADL_Overdrive6_PowerControlInfo_Get = (ADL_OVERDRIVE6_POWERCONTROLINFO_GET)GetProcAddress(hDLLA, "ADL_Overdrive6_PowerControlInfo_Get");
  ADL_Overdrive6_PowerControl_Get = (ADL_OVERDRIVE6_POWERCONTROL_GET)GetProcAddress(hDLLA, "ADL_Overdrive6_PowerControl_Get");
  ADL_Overdrive6_FanSpeed_Set  = (ADL_OVERDRIVE6_FANSPEED_SET)GetProcAddress(hDLLA, "ADL_Overdrive6_FanSpeed_Set");
  ADL_Overdrive6_State_Set = (ADL_OVERDRIVE6_STATE_SET)GetProcAddress(hDLLA, "ADL_Overdrive6_State_Set");
  ADL_Overdrive6_PowerControl_Set = (ADL_OVERDRIVE6_POWERCONTROL_SET) GetProcAddress(hDLLA, "ADL_Overdrive6_PowerControl_Set");

  if (!ADL_Overdrive6_FanSpeed_Get || !ADL_Overdrive6_ThermalController_Caps ||
    !ADL_Overdrive6_Temperature_Get || !ADL_Overdrive6_Capabilities_Get ||
    !ADL_Overdrive6_StateInfo_Get || !ADL_Overdrive6_CurrentStatus_Get ||
    !ADL_Overdrive6_PowerControl_Caps || !ADL_Overdrive6_PowerControlInfo_Get ||
    !ADL_Overdrive6_PowerControl_Get || !ADL_Overdrive6_FanSpeed_Set ||
    !ADL_Overdrive6_State_Set || !ADL_Overdrive6_PowerControl_Set) {
      PRINTFLN("ATI ADL Overdrive6's API is missing or broken.");
    return false;
  } else {
    PRINTFLN("ATI ADL Overdrive6 API found.");
  }

  return true;
}

bool prepare_adl(void)
{
  int result;

#if defined (__linux__)
  hDLLA = dlopen( "libatiadlxx.so", RTLD_LAZY|RTLD_GLOBAL);
#else
  hDLLA = LoadLibrary("atiadlxx.dll");
  if (hDLLA == NULL)
    // A 32 bit calling application on 64 bit OS will fail to LoadLIbrary.
    // Try to load the 32 bit library (atiadlxy.dll) instead
    hDLLA = LoadLibrary("atiadlxy.dll");
#endif
  if (hDLLA == NULL) {
    PRINTFLN("Unable to load ATI ADL library.");
    return false;
  }
  ADL_Main_Control_Create = (ADL_MAIN_CONTROL_CREATE) GetProcAddress(hDLLA,"ADL_Main_Control_Create");
  ADL_Main_Control_Destroy = (ADL_MAIN_CONTROL_DESTROY) GetProcAddress(hDLLA,"ADL_Main_Control_Destroy");
  ADL_Adapter_NumberOfAdapters_Get = (ADL_ADAPTER_NUMBEROFADAPTERS_GET) GetProcAddress(hDLLA,"ADL_Adapter_NumberOfAdapters_Get");
  ADL_Adapter_AdapterInfo_Get = (ADL_ADAPTER_ADAPTERINFO_GET) GetProcAddress(hDLLA,"ADL_Adapter_AdapterInfo_Get");
  ADL_Display_DisplayInfo_Get = (ADL_DISPLAY_DISPLAYINFO_GET) GetProcAddress(hDLLA,"ADL_Display_DisplayInfo_Get");
  ADL_Adapter_ID_Get = (ADL_ADAPTER_ID_GET) GetProcAddress(hDLLA,"ADL_Adapter_ID_Get");
  ADL_Main_Control_Refresh = (ADL_MAIN_CONTROL_REFRESH) GetProcAddress(hDLLA, "ADL_Main_Control_Refresh");
  ADL_Adapter_VideoBiosInfo_Get = (ADL_ADAPTER_VIDEOBIOSINFO_GET)GetProcAddress(hDLLA,"ADL_Adapter_VideoBiosInfo_Get");
  ADL_Overdrive_Caps = (ADL_OVERDRIVE_CAPS)GetProcAddress(hDLLA, "ADL_Overdrive_Caps");

  ADL_Adapter_Accessibility_Get = (ADL_ADAPTER_ACCESSIBILITY_GET)GetProcAddress(hDLLA, "ADL_Adapter_Accessibility_Get");

  if (!ADL_Main_Control_Create || !ADL_Main_Control_Destroy ||
    !ADL_Adapter_NumberOfAdapters_Get || !ADL_Adapter_AdapterInfo_Get ||
    !ADL_Display_DisplayInfo_Get ||
    !ADL_Adapter_ID_Get || !ADL_Main_Control_Refresh ||
    !ADL_Adapter_VideoBiosInfo_Get || !ADL_Overdrive_Caps) {
      PRINTFLN("ATI ADL API is missing or broken.");
    return false;
  }

  // Initialise ADL. The second parameter is 1, which means:
  // retrieve adapter information only for adapters that are physically present and enabled in the system
  result = ADL_Main_Control_Create(ADL_Main_Memory_Alloc, 1);
  if (result != ADL_OK) {
    PRINTFLN("ADL initialisation error: %d (%s)", result, adl_error_desc(result));
    return false;
  }

  result = ADL_Main_Control_Refresh();
  if (result != ADL_OK) {
    PRINTFLN("ADL refresh error: %d (%s)", result, adl_error_desc(result));
    return false;
  }

  init_overdrive5();
  init_overdrive6(); // FIXME: don't if ADL6 is not present

  return true;
}



float __gpu_temp(struct gpu_adl *ga)
{
  if (ADL_Overdrive5_Temperature_Get(ga->iAdapterIndex, 0, &ga->lpTemperature) != ADL_OK)
    return -1;
  return (float)ga->lpTemperature.iTemperature / 1000;
}

float gpu_temp(int gpu)
{
  struct gpu_adl *ga;
  float ret = -1;

  if (!gpus[gpu].has_adl || !adl_active)
    return ret;

  ga = &gpus[gpu].adl;
  lock_adl();
  ret = __gpu_temp(ga);
  unlock_adl();
  gpus[gpu].temp = ret;
  return ret;
}

inline int __gpu_engineclock(struct gpu_adl *ga)
{
  return ga->lpActivity.iEngineClock / 100;
}

int gpu_engineclock(int gpu)
{
  struct gpu_adl *ga;
  int ret = -1;

  if (!gpus[gpu].has_adl || !adl_active)
    return ret;

  ga = &gpus[gpu].adl;
  lock_adl();
  if (ADL_Overdrive5_CurrentActivity_Get(ga->iAdapterIndex, &ga->lpActivity) != ADL_OK)
    goto out;
  ret = __gpu_engineclock(ga);
out:
  unlock_adl();
  return ret;
}

inline int __gpu_memclock(struct gpu_adl *ga)
{
  return ga->lpActivity.iMemoryClock / 100;
}

int gpu_memclock(int gpu)
{
  struct gpu_adl *ga;
  int ret = -1;

  if (!gpus[gpu].has_adl || !adl_active)
    return ret;

  ga = &gpus[gpu].adl;
  lock_adl();
  if (ADL_Overdrive5_CurrentActivity_Get(ga->iAdapterIndex, &ga->lpActivity) != ADL_OK)
    goto out;
  ret = __gpu_memclock(ga);
out:
  unlock_adl();
  return ret;
}

inline float __gpu_vddc(struct gpu_adl *ga)
{
  printf("VDDC: %i\n", ga->lpActivity.iVddc);
  return (float)ga->lpActivity.iVddc / 1000;
}

float gpu_vddc(int gpu)
{
  struct gpu_adl *ga;
  float ret = -1;

  if (!gpus[gpu].has_adl || !adl_active)
    return ret;

  ga = &gpus[gpu].adl;
  lock_adl();
  if (ADL_Overdrive5_CurrentActivity_Get(ga->iAdapterIndex, &ga->lpActivity) != ADL_OK)
    goto out;
  ret = __gpu_vddc(ga);
out:
  unlock_adl();
  return ret;
}

inline int __gpu_activity(struct gpu_adl *ga)
{
  if (!ga->lpOdParameters.iActivityReportingSupported)
    return -1;
  return ga->lpActivity.iActivityPercent;
}

int gpu_activity(int gpu)
{
  struct gpu_adl *ga;
  int ret = -1;

  if (!gpus[gpu].has_adl || !adl_active)
    return ret;

  ga = &gpus[gpu].adl;
  lock_adl();
  ret = ADL_Overdrive5_CurrentActivity_Get(ga->iAdapterIndex, &ga->lpActivity);
  unlock_adl();
  if (ret != ADL_OK)
    return ret;
  if (!ga->lpOdParameters.iActivityReportingSupported)
    return ret;
  return ga->lpActivity.iActivityPercent;
}

inline int __gpu_fanspeed(struct gpu_adl *ga)
{
  if (!ga->has_fanspeed && ga->twin)
    return __gpu_fanspeed(ga->twin);

  if (!(ga->lpFanSpeedInfo.iFlags & ADL_DL_FANCTRL_SUPPORTS_RPM_READ))
    return -1;
  ga->lpFanSpeedValue.iSpeedType = ADL_DL_FANCTRL_SPEED_TYPE_RPM;
  if (ADL_Overdrive5_FanSpeed_Get(ga->iAdapterIndex, 0, &ga->lpFanSpeedValue) != ADL_OK)
    return -1;
  return ga->lpFanSpeedValue.iFanSpeed;
}

int gpu_fanspeed(int gpu)
{
  struct gpu_adl *ga;
  int ret = -1;

  if (!gpus[gpu].has_adl || !adl_active)
    return ret;

  ga = &gpus[gpu].adl;
  lock_adl();
  ret = __gpu_fanspeed(ga);
  unlock_adl();
  return ret;
}

int __gpu_fanpercent(struct gpu_adl *ga)
{
  if (!ga->has_fanspeed && ga->twin)
    return __gpu_fanpercent(ga->twin);

  if (!(ga->lpFanSpeedInfo.iFlags & ADL_DL_FANCTRL_SUPPORTS_PERCENT_READ ))
    return -1;
  ga->lpFanSpeedValue.iSpeedType = ADL_DL_FANCTRL_SPEED_TYPE_PERCENT;
  if (ADL_Overdrive5_FanSpeed_Get(ga->iAdapterIndex, 0, &ga->lpFanSpeedValue) != ADL_OK)
    return -1;
  return ga->lpFanSpeedValue.iFanSpeed;
}

float gpu_fanpercent(int gpu)
{
  struct gpu_adl *ga;
  int ret = -1;

  if (!gpus[gpu].has_adl || !adl_active)
    return ret;

  ga = &gpus[gpu].adl;
  lock_adl();
  ret = __gpu_fanpercent(ga);
  unlock_adl();
  return ret;
}

inline int __gpu_powertune(struct gpu_adl *ga)
{
  int dummy = 0;

  if (ADL_Overdrive5_PowerControl_Get(ga->iAdapterIndex, &ga->iPercentage, &dummy) != ADL_OK)
    return -1;
  return ga->iPercentage;
}

int gpu_powertune(int gpu)
{
  struct gpu_adl *ga;
  int ret = -1;

  if (!gpus[gpu].has_adl || !adl_active)
    return ret;

  ga = &gpus[gpu].adl;
  lock_adl();
  ret = __gpu_powertune(ga);
  unlock_adl();
  return ret;
}

bool gpu_stats(int gpu, float *temp, int *engineclock, int *memclock, float *vddc,
         int *activity, int *fanspeed, int *fanpercent, int *powertune)
{
  struct gpu_adl *ga;

  if (!gpus[gpu].has_adl || !adl_active)
    return false;

  ga = &gpus[gpu].adl;

  lock_adl();
  *temp = __gpu_temp(ga);
  if (ADL_Overdrive5_CurrentActivity_Get(ga->iAdapterIndex, &ga->lpActivity) != ADL_OK) {
    *engineclock = 0;
    *memclock = 0;
    *vddc = 0;
    *activity = 0;
  } else {
    *engineclock = __gpu_engineclock(ga);
    *memclock = __gpu_memclock(ga);
    *vddc = __gpu_vddc(ga);
    *activity = __gpu_activity(ga);
  }
  *fanspeed = __gpu_fanspeed(ga);
  *fanpercent = __gpu_fanpercent(ga);
  *powertune = __gpu_powertune(ga);
  unlock_adl();

  return true;
}

int set_memoryclock(int gpu, int iMemoryClock)
{
  ADLODPerformanceLevels *lpOdPerformanceLevels;
  int i, lev, ret = 1;
  struct gpu_adl *ga;

  if (!gpus[gpu].has_adl || !adl_active) {
    PRINTFLN("Set memoryclock not supported\n");
    return ret;
  }

  iMemoryClock *= 100;
  ga = &gpus[gpu].adl;

  lev = ga->lpOdParameters.iNumberOfPerformanceLevels - 1;
  lpOdPerformanceLevels = (ADLODPerformanceLevels *)alloca(sizeof(ADLODPerformanceLevels)+(lev * sizeof(ADLODPerformanceLevel)));
  lpOdPerformanceLevels->iSize = sizeof(ADLODPerformanceLevels) + sizeof(ADLODPerformanceLevel) * lev;

  lock_adl();
  if (ADL_Overdrive5_ODPerformanceLevels_Get(ga->iAdapterIndex, 0, lpOdPerformanceLevels) != ADL_OK)
    goto out;
  lpOdPerformanceLevels->aLevels[lev].iMemoryClock = iMemoryClock;
  for (i = 0; i < lev; i++) {
    if (lpOdPerformanceLevels->aLevels[i].iMemoryClock > iMemoryClock)
      lpOdPerformanceLevels->aLevels[i].iMemoryClock = iMemoryClock;
  }
  ADL_Overdrive5_ODPerformanceLevels_Set(ga->iAdapterIndex, lpOdPerformanceLevels);
  ADL_Overdrive5_ODPerformanceLevels_Get(ga->iAdapterIndex, 0, lpOdPerformanceLevels);
  if (lpOdPerformanceLevels->aLevels[lev].iMemoryClock == iMemoryClock)
    ret = 0;
  ga->iEngineClock = lpOdPerformanceLevels->aLevels[lev].iEngineClock;
  ga->iMemoryClock = lpOdPerformanceLevels->aLevels[lev].iMemoryClock;
  ga->iVddc = lpOdPerformanceLevels->aLevels[lev].iVddc;
  ga->managed = true;
out:
  unlock_adl();
  return ret;
}

int set_engineclock(int gpu, int iEngineClock)
{
  ADLODPerformanceLevels *lpOdPerformanceLevels;
  struct cgpu_info *cgpu;
  int i, lev, ret = 1;
  struct gpu_adl *ga;

  if (!gpus[gpu].has_adl || !adl_active) {
    PRINTFLN("Set engineclock not supported\n");
    return ret;
  }

  iEngineClock *= 100;
  ga = &gpus[gpu].adl;

  /* Keep track of intended engine clock in case the device changes
   * profile and drops while idle, not taking the new engine clock */
  ga->lastengine = iEngineClock;

  lev = ga->lpOdParameters.iNumberOfPerformanceLevels - 1;
  lpOdPerformanceLevels = (ADLODPerformanceLevels *)alloca(sizeof(ADLODPerformanceLevels)+(lev * sizeof(ADLODPerformanceLevel)));
  lpOdPerformanceLevels->iSize = sizeof(ADLODPerformanceLevels) + sizeof(ADLODPerformanceLevel) * lev;

  lock_adl();
  if (ADL_Overdrive5_ODPerformanceLevels_Get(ga->iAdapterIndex, 0, lpOdPerformanceLevels) != ADL_OK)
    goto out;
  for (i = 0; i < lev; i++) {
    if (lpOdPerformanceLevels->aLevels[i].iEngineClock > iEngineClock)
      lpOdPerformanceLevels->aLevels[i].iEngineClock = iEngineClock;
  }
  lpOdPerformanceLevels->aLevels[lev].iEngineClock = iEngineClock;
  ADL_Overdrive5_ODPerformanceLevels_Set(ga->iAdapterIndex, lpOdPerformanceLevels);
  ADL_Overdrive5_ODPerformanceLevels_Get(ga->iAdapterIndex, 0, lpOdPerformanceLevels);
  if (lpOdPerformanceLevels->aLevels[lev].iEngineClock == iEngineClock)
    ret = 0;
  ga->iEngineClock = lpOdPerformanceLevels->aLevels[lev].iEngineClock;
  if (ga->iEngineClock > ga->maxspeed)
    ga->maxspeed = ga->iEngineClock;
  if (ga->iEngineClock < ga->minspeed)
    ga->minspeed = ga->iEngineClock;
  ga->iMemoryClock = lpOdPerformanceLevels->aLevels[lev].iMemoryClock;
  ga->iVddc = lpOdPerformanceLevels->aLevels[lev].iVddc;
  ga->managed = true;
out:
  unlock_adl();

  cgpu = &gpus[gpu];
  if (cgpu->gpu_memdiff)
    set_memoryclock(gpu, iEngineClock / 100 + cgpu->gpu_memdiff);

  return ret;
}

int set_vddc(int gpu, float fVddc)
{
  ADLODPerformanceLevels *lpOdPerformanceLevels;
  int i, iVddc, lev, ret = 1;
  struct gpu_adl *ga;

  if (!gpus[gpu].has_adl || !adl_active) {
    PRINTFLN("Set vddc not supported\n");
    return ret;
  }

  iVddc = 1000 * fVddc;
  ga = &gpus[gpu].adl;

  lev = ga->lpOdParameters.iNumberOfPerformanceLevels - 1;
  lpOdPerformanceLevels = (ADLODPerformanceLevels *)alloca(sizeof(ADLODPerformanceLevels)+(lev * sizeof(ADLODPerformanceLevel)));
  lpOdPerformanceLevels->iSize = sizeof(ADLODPerformanceLevels) + sizeof(ADLODPerformanceLevel) * lev;

  lock_adl();
  if (ADL_Overdrive5_ODPerformanceLevels_Get(ga->iAdapterIndex, 0, lpOdPerformanceLevels) != ADL_OK)
    goto out;
  for (i = 0; i < lev; i++) {
    if (lpOdPerformanceLevels->aLevels[i].iVddc > iVddc)
      lpOdPerformanceLevels->aLevels[i].iVddc = iVddc;
  }
  lpOdPerformanceLevels->aLevels[lev].iVddc = iVddc;
  ADL_Overdrive5_ODPerformanceLevels_Set(ga->iAdapterIndex, lpOdPerformanceLevels);
  ADL_Overdrive5_ODPerformanceLevels_Get(ga->iAdapterIndex, 0, lpOdPerformanceLevels);
  if (lpOdPerformanceLevels->aLevels[lev].iVddc == iVddc)
    ret = 0;
  ga->iEngineClock = lpOdPerformanceLevels->aLevels[lev].iEngineClock;
  ga->iMemoryClock = lpOdPerformanceLevels->aLevels[lev].iMemoryClock;
  ga->iVddc = lpOdPerformanceLevels->aLevels[lev].iVddc;
  ga->managed = true;
out:
  unlock_adl();
  return ret;
}

void get_fanrange(int gpu, int *imin, int *imax)
{
  struct gpu_adl *ga;

  if (!gpus[gpu].has_adl || !adl_active) {
    PRINTFLN("Get fanrange not supported\n");
    return;
  }
  ga = &gpus[gpu].adl;
  *imin = ga->lpFanSpeedInfo.iMinPercent;
  *imax = ga->lpFanSpeedInfo.iMaxPercent;
}

int set_fanspeed(int gpu, float FanSpeed)
{
  struct gpu_adl *ga;
  int ret = 1;
  int iFanSpeed = FanSpeed;

  if (!gpus[gpu].has_adl || !adl_active) {
    PRINTFLN("Set fanspeed not supported\n");
    return ret;
  }

  ga = &gpus[gpu].adl;
  if (!(ga->lpFanSpeedInfo.iFlags & (ADL_DL_FANCTRL_SUPPORTS_RPM_WRITE | ADL_DL_FANCTRL_SUPPORTS_PERCENT_WRITE ))) {
    PRINTFLN("GPU %d doesn't support rpm or percent write", gpu);
    return ret;
  }

  /* Store what fanspeed we're actually aiming for for re-entrant changes
   * in case this device does not support fine setting changes */
  ga->targetfan = iFanSpeed;

  lock_adl();
  ga->lpFanSpeedValue.iSpeedType = ADL_DL_FANCTRL_SPEED_TYPE_PERCENT;
  if (ADL_Overdrive5_FanSpeed_Get(ga->iAdapterIndex, 0, &ga->lpFanSpeedValue) != ADL_OK) {
    PRINTFLN("GPU %d call to fanspeed get failed", gpu);
  }
  if (!(ga->lpFanSpeedValue.iFlags & ADL_DL_FANCTRL_FLAG_USER_DEFINED_SPEED)) {
    /* If user defined is not already specified, set it first */
    ga->lpFanSpeedValue.iFlags |= ADL_DL_FANCTRL_FLAG_USER_DEFINED_SPEED;
    ADL_Overdrive5_FanSpeed_Set(ga->iAdapterIndex, 0, &ga->lpFanSpeedValue);
  }
  if (!(ga->lpFanSpeedInfo.iFlags & ADL_DL_FANCTRL_SUPPORTS_PERCENT_WRITE)) {
    /* Must convert speed to an RPM */
    iFanSpeed = ga->lpFanSpeedInfo.iMaxRPM * iFanSpeed / 100;
    ga->lpFanSpeedValue.iSpeedType = ADL_DL_FANCTRL_SPEED_TYPE_RPM;
  } else
    ga->lpFanSpeedValue.iSpeedType = ADL_DL_FANCTRL_SPEED_TYPE_PERCENT;
  ga->lpFanSpeedValue.iFanSpeed = iFanSpeed;
  ret = ADL_Overdrive5_FanSpeed_Set(ga->iAdapterIndex, 0, &ga->lpFanSpeedValue);
  ga->managed = true;
  unlock_adl();

  return ret;
}

int set_powertune(int gpu, int iPercentage)
{
  struct gpu_adl *ga;
  int dummy, ret = 1;

  if (!gpus[gpu].has_adl || !adl_active) {
    PRINTFLN("Set powertune not supported\n");
    return ret;
  }

  ga = &gpus[gpu].adl;

  lock_adl();
  ADL_Overdrive5_PowerControl_Set(ga->iAdapterIndex, iPercentage);
  ADL_Overdrive5_PowerControl_Get(ga->iAdapterIndex, &ga->iPercentage, &dummy);
  if (ga->iPercentage == iPercentage)
    ret = 0;
  ga->managed = true;
  unlock_adl();
  return ret;
}

void set_defaultfan(int gpu)
{
  struct gpu_adl *ga;
  if (!gpus[gpu].has_adl || !adl_active)
    return;

  ga = &gpus[gpu].adl;
  lock_adl();
  if (ga->def_fan_valid)
    ADL_Overdrive5_FanSpeed_Set(ga->iAdapterIndex, 0, &ga->DefFanSpeedValue);
  unlock_adl();
}

void set_defaultengine(int gpu)
{
  struct gpu_adl *ga;
  if (!gpus[gpu].has_adl || !adl_active)
    return;

  ga = &gpus[gpu].adl;
  lock_adl();
  if (ga->DefPerfLev)
    ADL_Overdrive5_ODPerformanceLevels_Set(ga->iAdapterIndex, ga->DefPerfLev);
  unlock_adl();
}

void init_adl(int nDevs)
{
  int result, i, j, devices = 0, last_adapter = -1, gpu = 0, dummy = 0;
  struct gpu_adapters adapters[MAX_GPUDEVICES], vadapters[MAX_GPUDEVICES];
  bool devs_match = true;
  ADLBiosInfo BiosInfo;

  PRINTFLN("Number of ADL devices: %d", nDevs);

  if (unlikely(pthread_mutex_init(&adl_lock, NULL))) {
    PRINTFLN("Failed to init adl_lock in init_adl");
    return;
  }

  if (!prepare_adl())
    return;

  // Obtain the number of adapters for the system
  result = ADL_Adapter_NumberOfAdapters_Get (&iNumberAdapters);
  if (result != ADL_OK) {
    PRINTFLN("Cannot get the number of adapters! Error %d!", result);
    return ;
  }

  if (iNumberAdapters > 0) {
    lpInfo = (LPAdapterInfo)malloc ( sizeof (AdapterInfo) * iNumberAdapters );
    memset ( lpInfo,'\0', sizeof (AdapterInfo) * iNumberAdapters );

    lpInfo->iSize = sizeof(lpInfo);
    // Get the AdapterInfo structure for all adapters in the system
    result = ADL_Adapter_AdapterInfo_Get (lpInfo, sizeof (AdapterInfo) * iNumberAdapters);
    if (result != ADL_OK) {
      PRINTFLN("ADL_Adapter_AdapterInfo_Get Error! Error %d", result);
      return ;
    }
  } else {
    PRINTFLN("No adapters found");
    return;
  }

  PRINTFLN("Found %d logical ADL adapters", iNumberAdapters);

  /* Iterate over iNumberAdapters and find the lpAdapterID of real devices */
  for (i = 0; i < iNumberAdapters; i++) {
    int iAdapterIndex;
    int lpAdapterID;

    iAdapterIndex = lpInfo[i].iAdapterIndex;

    /* Get unique identifier of the adapter, 0 means not AMD */
    result = ADL_Adapter_ID_Get(iAdapterIndex, &lpAdapterID);

    if (ADL_Adapter_VideoBiosInfo_Get(iAdapterIndex, &BiosInfo) == ADL_ERR) {
      //PRINTFLN("ADL index %d, id %d - FAILED to get BIOS info", iAdapterIndex, lpAdapterID);
    } else {
      //PRINTFLN("ADL index %d, id %d - BIOS partno.: %s, version: %s, date: %s", iAdapterIndex, lpAdapterID, BiosInfo.strPartNumber, BiosInfo.strVersion, BiosInfo.strDate);
    }

    if (result != ADL_OK) {
      //PRINTFLN("Failed to ADL_Adapter_ID_Get. Error %d", result);
      if (result == -10)
        PRINTFLN("(Device is not enabled.)");
      continue;
    }

    /* Each adapter may have multiple entries */
    if (lpAdapterID == last_adapter) {
      continue;
    }

    PRINTFLN("GPU %d assigned: "
           "iAdapterIndex:%d "
           "iPresent:%d "
           "strUDID:%s "
           "iBusNumber:%d "
           "iDeviceNumber:%d "
#if defined(__linux__)
           "iDrvIndex:%d "
#endif
           "iFunctionNumber:%d "
           "iVendorID:%d "
           "name:%s",
           devices,
           lpInfo[i].iAdapterIndex,
           lpInfo[i].iPresent,
           lpInfo[i].strUDID,
           lpInfo[i].iBusNumber,
           lpInfo[i].iDeviceNumber,
#if defined(__linux__)
           lpInfo[i].iDrvIndex,
#endif
           lpInfo[i].iFunctionNumber,
           lpInfo[i].iVendorID,
           lpInfo[i].strAdapterName);

    adapters[devices].iAdapterIndex = iAdapterIndex;
    adapters[devices].iBusNumber = lpInfo[i].iBusNumber;
    adapters[devices].id = i;

    /* We found a truly new adapter instead of a logical
     * one. Now since there's no way of correlating the
     * opencl enumerated devices and the ADL enumerated
     * ones, we have to assume they're in the same order.*/
    if (++devices > nDevs && devs_match) {
      PRINTFLN("ADL found more devices than opencl!");
      PRINTFLN("There is possibly at least one GPU that doesn't support OpenCL");
      PRINTFLN("Use the gpu map feature to reliably map OpenCL to ADL");
      devs_match = false;
    }
    last_adapter = lpAdapterID;

    if (!lpAdapterID) {
      PRINTFLN("Adapter returns ID 0 meaning not AMD. Card order might be confused");
      continue;
    }
  }

  if (devices < nDevs) {
    PRINTFLN("ADL found less devices than opencl!");
    PRINTFLN("There is possibly more than one display attached to a GPU");
    PRINTFLN("Use the gpu map feature to reliably map OpenCL to ADL");
    devs_match = false;
  }

  for (i = 0; i < devices; i++) {
    vadapters[i].virtual_gpu = i;
    vadapters[i].id = adapters[i].id;
  }

  /* Apply manually provided OpenCL to ADL mapping, if any */
  for (i = 0; i < nDevs; i++) {
    if (gpus[i].mapped) {
      vadapters[gpus[i].virtual_adl].virtual_gpu = i;
      PRINTFLN("Mapping OpenCL device %d to ADL device %d", i, gpus[i].virtual_adl);
    } else {
      gpus[i].virtual_adl = i;
    }
  }

  if (!devs_match) {
    PRINTFLN("WARNING: Number of OpenCL and ADL devices did not match!");
    PRINTFLN("Hardware monitoring may NOT match up with devices!");
  } else {
    /* Windows has some kind of random ordering for bus number IDs and
     * ordering the GPUs according to ascending order fixes it. Linux
     * has usually sequential but decreasing order instead! */
    for (i = 0; i < devices; i++) {
      int j, virtual_gpu;

      virtual_gpu = 0;
      for (j = 0; j < devices; j++) {
        if (i == j)
          continue;
#ifdef WIN32
        if (adapters[j].iBusNumber < adapters[i].iBusNumber)
#else
        if (adapters[j].iBusNumber > adapters[i].iBusNumber)
#endif
          virtual_gpu++;
      }
      if (virtual_gpu != i) {
        PRINTFLN("Mapping device %d to GPU %d according to Bus Number order",
               i, virtual_gpu);
        vadapters[virtual_gpu].virtual_gpu = i;
        vadapters[virtual_gpu].id = adapters[i].id;
      }
    }
  }

  if (devices > nDevs)
    devices = nDevs;

  for (gpu = 0; gpu < devices; gpu++) {
    struct gpu_adl *ga;
    int iAdapterIndex;
    int lpAdapterID;
    ADLODPerformanceLevels *lpOdPerformanceLevels;
    int lev, adlGpu;
    size_t plsize;
    ADLBiosInfo BiosInfo;

    adlGpu = gpus[gpu].virtual_adl;
    i = vadapters[adlGpu].id;
    iAdapterIndex = lpInfo[i].iAdapterIndex;
    gpus[gpu].virtual_gpu = vadapters[adlGpu].virtual_gpu;

    /* Get unique identifier of the adapter, 0 means not AMD */
    result = ADL_Adapter_ID_Get(iAdapterIndex, &lpAdapterID);
    if (result != ADL_OK) {
      PRINTFLN("Failed to ADL_Adapter_ID_Get. Error %d", result);
      continue;
    }

    PRINTFLN("GPU %d %s hardware monitoring enabled", gpu, lpInfo[i].strAdapterName);
    if (gpus[gpu].name)
      free(gpus[gpu].name);
    gpus[gpu].name = lpInfo[i].strAdapterName;
    gpus[gpu].has_adl = true;
    /* Flag adl as active if any card is successfully activated */
    adl_active = true;

    /* From here on we know this device is a discrete device and
     * should support ADL */
    ga = &gpus[gpu].adl;
    ga->gpu = gpu;
    ga->iAdapterIndex = iAdapterIndex;
    ga->lpAdapterID = lpAdapterID;
    strcpy(ga->strAdapterName, lpInfo[i].strAdapterName);
    ga->DefPerfLev = NULL;
    ga->twin = NULL;
    ga->def_fan_valid = false;

    PRINTFLN("ADL GPU %d is Adapter index %d and maps to adapter id %d", ga->gpu, ga->iAdapterIndex, ga->lpAdapterID);

    if (ADL_Adapter_VideoBiosInfo_Get(iAdapterIndex, &BiosInfo) != ADL_ERR)
      PRINTFLN("GPU %d BIOS partno.: %s, version: %s, date: %s", gpu, BiosInfo.strPartNumber, BiosInfo.strVersion, BiosInfo.strDate);

    ga->lpOdParameters.iSize = sizeof(ADLODParameters);
    if (ADL_Overdrive5_ODParameters_Get(iAdapterIndex, &ga->lpOdParameters) != ADL_OK)
      PRINTFLN("Failed to ADL_Overdrive5_ODParameters_Get");

    lev = ga->lpOdParameters.iNumberOfPerformanceLevels - 1;
    /* We're only interested in the top performance level */
    plsize = sizeof(ADLODPerformanceLevels) + lev * sizeof(ADLODPerformanceLevel);
    lpOdPerformanceLevels = (ADLODPerformanceLevels *)malloc(plsize);
    lpOdPerformanceLevels->iSize = plsize;

    /* Get default performance levels first */
    //if (ADL_Overdrive5_ODPerformanceLevels_Get(iAdapterIndex, 1, lpOdPerformanceLevels) != ADL_OK)
      //PRINTFLN("Failed to get default ADL_Overdrive5_ODPerformanceLevels_Get");
    /* Set the limits we'd use based on default gpu speeds */
    ga->maxspeed = ga->minspeed = lpOdPerformanceLevels->aLevels[lev].iEngineClock;

    ga->lpTemperature.iSize = sizeof(ADLTemperature);
    ga->lpFanSpeedInfo.iSize = sizeof(ADLFanSpeedInfo);
    ga->lpFanSpeedValue.iSize = ga->DefFanSpeedValue.iSize = sizeof(ADLFanSpeedValue);
    ga->lpFanSpeedValue.iSpeedType = ADL_DL_FANCTRL_SPEED_TYPE_RPM;
    ga->DefFanSpeedValue.iSpeedType = ADL_DL_FANCTRL_SPEED_TYPE_RPM;

    /* Now get the current performance levels for any existing overclock */
    if (ADL_Overdrive5_ODPerformanceLevels_Get(iAdapterIndex, 0, lpOdPerformanceLevels) != ADL_OK)
      PRINTFLN("Failed to get exist ADL_Overdrive5_ODPerformanceLevels_Get");
    else {
      /* Save these values as the defaults in case we wish to reset to defaults */
      ga->DefPerfLev = (ADLODPerformanceLevels *)malloc(plsize);
      memcpy(ga->DefPerfLev, lpOdPerformanceLevels, plsize);
    }

    if (gpus[gpu].gpu_engine) {
      int setengine = gpus[gpu].gpu_engine * 100;

      /* Lower profiles can't have a higher setting */
      for (j = 0; j < lev; j++) {
        if (lpOdPerformanceLevels->aLevels[j].iEngineClock > setengine)
          lpOdPerformanceLevels->aLevels[j].iEngineClock = setengine;
      }
      lpOdPerformanceLevels->aLevels[lev].iEngineClock = setengine;
      PRINTFLN("Setting GPU %d engine clock to %d", gpu, gpus[gpu].gpu_engine);
      ADL_Overdrive5_ODPerformanceLevels_Set(iAdapterIndex, lpOdPerformanceLevels);
      ga->maxspeed = setengine;
      if (gpus[gpu].min_engine)
        ga->minspeed = gpus[gpu].min_engine * 100;
      ga->managed = true;
      if (gpus[gpu].gpu_memdiff)
        set_memoryclock(gpu, gpus[gpu].gpu_engine + gpus[gpu].gpu_memdiff);
    }

    if (gpus[gpu].gpu_memclock) {
      int setmem = gpus[gpu].gpu_memclock * 100;

      for (j = 0; j < lev; j++) {
        if (lpOdPerformanceLevels->aLevels[j].iMemoryClock > setmem)
          lpOdPerformanceLevels->aLevels[j].iMemoryClock = setmem;
      }
      lpOdPerformanceLevels->aLevels[lev].iMemoryClock = setmem;
      PRINTFLN("Setting GPU %d memory clock to %d", gpu, gpus[gpu].gpu_memclock);
      ADL_Overdrive5_ODPerformanceLevels_Set(iAdapterIndex, lpOdPerformanceLevels);
      ga->managed = true;
    }

    if (gpus[gpu].gpu_vddc) {
      int setv = gpus[gpu].gpu_vddc * 1000;

      for (j = 0; j < lev; j++) {
        if (lpOdPerformanceLevels->aLevels[j].iVddc > setv)
          lpOdPerformanceLevels->aLevels[j].iVddc = setv;
      }
      lpOdPerformanceLevels->aLevels[lev].iVddc = setv;
      PRINTFLN("Setting GPU %d voltage to %.3f", gpu, gpus[gpu].gpu_vddc);
      ADL_Overdrive5_ODPerformanceLevels_Set(iAdapterIndex, lpOdPerformanceLevels);
      ga->managed = true;
    }

    ADL_Overdrive5_ODPerformanceLevels_Get(iAdapterIndex, 0, lpOdPerformanceLevels);
    ga->iEngineClock = lpOdPerformanceLevels->aLevels[lev].iEngineClock;
    ga->iMemoryClock = lpOdPerformanceLevels->aLevels[lev].iMemoryClock;
    ga->iVddc = lpOdPerformanceLevels->aLevels[lev].iVddc;
    ga->iBusNumber = lpInfo[i].iBusNumber;

    if (ADL_Overdrive5_FanSpeedInfo_Get(iAdapterIndex, 0, &ga->lpFanSpeedInfo) != ADL_OK)
      PRINTFLN("Failed to ADL_Overdrive5_FanSpeedInfo_Get");

    if(!(ga->lpFanSpeedInfo.iFlags & (ADL_DL_FANCTRL_SUPPORTS_RPM_WRITE | ADL_DL_FANCTRL_SUPPORTS_PERCENT_WRITE)))
      ga->has_fanspeed = false;
    else
      ga->has_fanspeed = true;

    /* Save the fanspeed values as defaults in case we reset later */
    if (ADL_Overdrive5_FanSpeed_Get(iAdapterIndex, 0, &ga->DefFanSpeedValue) != ADL_OK)
      PRINTFLN("Failed to ADL_Overdrive5_FanSpeed_Get for default value");
    else
      ga->def_fan_valid = true;

    if (gpus[gpu].gpu_fan)
      set_fanspeed(gpu, gpus[gpu].gpu_fan);
    else
      gpus[gpu].gpu_fan = 85; /* Set a nominal upper limit of 85% */

    /* Not fatal if powercontrol get fails */
    if (ADL_Overdrive5_PowerControl_Get(iAdapterIndex, &ga->iPercentage, &dummy) != ADL_OK)
      PRINTFLN("Failed to ADL_Overdrive5_PowerControl_get");

    if (gpus[gpu].gpu_powertune) {
      ADL_Overdrive5_PowerControl_Set(iAdapterIndex, gpus[gpu].gpu_powertune);
      ADL_Overdrive5_PowerControl_Get(iAdapterIndex, &ga->iPercentage, &dummy);
      ga->managed = true;
    }

    ga->lasttemp = __gpu_temp(ga);
  }

  for (gpu = 0; gpu < devices; gpu++) {
    struct gpu_adl *ga = &gpus[gpu].adl;
    int j;

    for (j = 0; j < devices; j++) {
      struct gpu_adl *other_ga;

      if (j == gpu)
        continue;

      other_ga = &gpus[j].adl;

      /* Search for twin GPUs on a single card. They will be
       * separated by one bus id and one will have fanspeed
       * while the other won't. */
      if (!ga->has_fanspeed) {
        if (fanspeed_twin(ga, other_ga)) {
          PRINTFLN("Dual GPUs detected: %d and %d",
            ga->gpu, other_ga->gpu);
          ga->twin = other_ga;
          other_ga->twin = ga;
        }
      }
    }
  }
}

void free_adl(void)
{
  if (hDLLA == NULL) return;

  ADL_Main_Memory_Free ((void **)&lpInfo);
  ADL_Main_Control_Destroy ();
#if defined (__linux__)
  dlclose(hDLLA);
#else
  FreeLibrary(hDLLA);
#endif
}

void clear_adl(int nDevs)
{
  struct gpu_adl *ga;
  int i;

  if (!adl_active)
    return;

  lock_adl();
  /* Try to reset values to their defaults */
  for (i = 0; i < nDevs; i++) {
    ga = &gpus[i].adl;
    /*  Only reset the values if we've changed them at any time */
    if (!gpus[i].has_adl || !ga->managed || !ga->DefPerfLev)
      continue;
    ADL_Overdrive5_ODPerformanceLevels_Set(ga->iAdapterIndex, ga->DefPerfLev);
    free(ga->DefPerfLev);
    if (ga->def_fan_valid)
      ADL_Overdrive5_FanSpeed_Set(ga->iAdapterIndex, 0, &ga->DefFanSpeedValue);
    ADL_Overdrive5_FanSpeedToDefault_Set(ga->iAdapterIndex, 0);
  }
  adl_active = false;
  unlock_adl();
  free_adl();
}

#endif
