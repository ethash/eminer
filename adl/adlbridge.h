#include <stdio.h>
#include <string.h>
#include <stdbool.h>
#include <math.h>
#include <pthread.h>

#if (defined(__linux__) || defined (WIN32))

#define MAX_GPUDEVICES 16

#define DEBUG 0
#define PRINTFLN(...) do { if(DEBUG) { printf(__VA_ARGS__); puts(""); } } while (0)

#define unlikely(expr) (__builtin_expect(!!(expr), 0))
#define likely(expr) (__builtin_expect(!!(expr), 1))

# include <wchar.h>
# include "adl_sdk.h"

#if defined (__linux__)
# include <dlfcn.h>
# include <stdlib.h>
# include <unistd.h>
#else /* WIN32 */
# include <windows.h>
# include <tchar.h>
#endif
#include "adl_functions.h"

extern bool opt_reorder;
extern int opt_hysteresis;
extern int opt_targettemp;
extern int opt_overheattemp;

pthread_mutex_t adl_lock;

struct gpu_adapters {
  int iAdapterIndex;
  int iBusNumber;
  int virtual_gpu;
  int id;
};

struct gpu_adl {
  int iAdapterIndex;
  int lpAdapterID;
  int iBusNumber;
  char strAdapterName[256];

  ADLTemperature lpTemperature;
  ADLPMActivity lpActivity;
  ADLODParameters lpOdParameters;
  ADLODPerformanceLevels *DefPerfLev;
  ADLFanSpeedInfo lpFanSpeedInfo;
  ADLFanSpeedValue lpFanSpeedValue;
  ADLFanSpeedValue DefFanSpeedValue;

  bool def_fan_valid;

  int iEngineClock;
  int iMemoryClock;
  int iVddc;
  int iPercentage;

  bool autofan;
  bool autoengine;
  bool managed; /* Were the values ever changed on this card */

  int lastengine;
  int lasttemp;
  int targetfan;
  int targettemp;
  int overtemp;
  int minspeed;
  int maxspeed;

  int gpu;
  bool has_fanspeed;
  struct gpu_adl *twin;
};

struct cgpu_info {
  char *name;

  bool mapped;
  int virtual_gpu;
  int virtual_adl;

  bool has_adl;
  struct gpu_adl adl;

  int gpu_engine;
  int min_engine;
  int gpu_fan;
  int min_fan;
  int gpu_memclock;
  int gpu_memdiff;
  int gpu_powertune;
  float gpu_vddc;

  float temp;
};

struct cgpu_info gpus[MAX_GPUDEVICES];
double total_secs;

int gpu_engine;
int min_engine;
int gpu_fan;
int min_fan;
int gpu_memdiff;

char *adl_error_desc(int error);

void init_adl(int nDevs);
void free_adl();

int gpu_index(int bus_number);

float gpu_temp(int deviceid);
float gpu_fanpercent(int deviceid);

int gpu_memclock(int deviceid);
int gpu_engineclock(int deviceid);

int set_fanspeed(int deviceid, float fanPercent);
int set_engineclock(int deviceid, int value);
int set_memoryclock(int deviceid, int value);

char *gpu_name(int deviceid);

#endif
