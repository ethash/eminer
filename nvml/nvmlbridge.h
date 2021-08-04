#include <stdlib.h>
#include <stddef.h>
#include <string.h>
#include <stdio.h>
#include "nvml.h"

#define DEBUG 0
#define PRINTFLN(...) do { if(DEBUG) { printf(__VA_ARGS__); puts(""); } } while (0)

#if defined(__linux__) || defined(WIN32)

#ifdef __linux__
#include <unistd.h>
#include <dlfcn.h>

void *hDLLN;
#elif WIN32
#include <windows.h>
typedef unsigned int uint;

#define dlsym (void *) GetProcAddress
#define dlclose FreeLibrary

static HMODULE hDLLN;
#endif

nvmlReturn_t init_nvml();

char * bridge_nvmlErrorString(nvmlReturn_t);
nvmlReturn_t bridge_nvmlDeviceGetPciInfo(nvmlDevice_t d, nvmlPciInfo_t * p);
nvmlReturn_t bridge_nvmlDeviceGetIndex(nvmlDevice_t d, uint * i);
nvmlReturn_t bridge_nvmlDeviceGetHandleByIndex(uint i, nvmlDevice_t * d);
nvmlReturn_t bridge_nvmlDeviceGetUUID(nvmlDevice_t d, char * u, uint s);
nvmlReturn_t bridge_nvmlDeviceGetName(nvmlDevice_t d, char * n, uint s);
nvmlReturn_t bridge_nvmlDeviceGetCount(uint * c);
nvmlReturn_t bridge_nvmlDeviceGetTemperature(nvmlDevice_t d, nvmlTemperatureSensors_t t, uint * v);
nvmlReturn_t bridge_nvmlDeviceGetFanSpeed(nvmlDevice_t d, uint * v);
nvmlReturn_t bridge_nvmlDeviceGetClockInfo(nvmlDevice_t d, nvmlClockType_t t, uint * r);
nvmlReturn_t bridge_nvmlShutdown();

static char * (*NVML_nvmlErrorString)(nvmlReturn_t);
static nvmlReturn_t (*NVML_nvmlInit)();
static nvmlReturn_t (*NVML_nvmlDeviceGetCount)(uint *);
static nvmlReturn_t (*NVML_nvmlDeviceGetIndex)(nvmlDevice_t, uint *);
static nvmlReturn_t (*NVML_nvmlDeviceGetUUID)(nvmlDevice_t, char *, uint);
static nvmlReturn_t (*NVML_nvmlDeviceGetHandleByIndex)(uint, nvmlDevice_t *);
static nvmlReturn_t (*NVML_nvmlDeviceGetName)(nvmlDevice_t, char *, uint);
static nvmlReturn_t (*NVML_nvmlDeviceGetPciInfo)(nvmlDevice_t, nvmlPciInfo_t *);
static nvmlReturn_t (*NVML_nvmlDeviceGetTemperature)(nvmlDevice_t, nvmlTemperatureSensors_t, uint *);
static nvmlReturn_t (*NVML_nvmlDeviceGetFanSpeed)(nvmlDevice_t, uint *);
static nvmlReturn_t (*NVML_nvmlDeviceGetClockInfo)(nvmlDevice_t, nvmlClockType_t, uint *);
static nvmlReturn_t (*NVML_nvmlShutdown)();

// Not every function can be genericized in this way because of all the custom structs,
// but there are several nvmlGet functions we want that take a nvmlDevice_t, *char, and
// a length as arguments. These are trivial to pass as function pointers along with their,
// arguments, so we might as well save some effort.
typedef int (*gettextProperty) (nvmlDevice_t device , char *buf, unsigned int length);
int bridge_get_text_property(gettextProperty f,
                             nvmlDevice_t device,
                             char *buf,
                             unsigned int length);

// Same as above, but for integer properties
typedef int (*getintProperty) (nvmlDevice_t device , unsigned int *property);
int bridge_get_int_property(getintProperty f,
                             nvmlDevice_t device,
                             unsigned int *property);

#endif
