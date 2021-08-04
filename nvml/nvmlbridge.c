
#include "nvmlbridge.h"

#if defined(__linux__) || defined(WIN32)

char * bridge_nvmlErrorString(nvmlReturn_t r) {
    return NVML_nvmlErrorString(r);
}

nvmlReturn_t bridge_nvmlDeviceGetIndex(nvmlDevice_t d, uint * i) {
    return NVML_nvmlDeviceGetIndex(d, i);
}

nvmlReturn_t bridge_nvmlDeviceGetHandleByIndex(uint i, nvmlDevice_t * d) {
    return NVML_nvmlDeviceGetHandleByIndex(i, d);
}

nvmlReturn_t bridge_nvmlDeviceGetUUID(nvmlDevice_t d, char * u, uint s) {
    return NVML_nvmlDeviceGetUUID(d, u, s);
}

nvmlReturn_t bridge_nvmlDeviceGetName(nvmlDevice_t d, char * n, uint s) {
    return NVML_nvmlDeviceGetName(d, n, s);
}

nvmlReturn_t bridge_nvmlDeviceGetTemperature(nvmlDevice_t d, nvmlTemperatureSensors_t t, uint * v) {
    return NVML_nvmlDeviceGetTemperature(d, t, v);
}

nvmlReturn_t bridge_nvmlDeviceGetFanSpeed(nvmlDevice_t d, uint * v) {
    return NVML_nvmlDeviceGetFanSpeed(d, v);
}

nvmlReturn_t bridge_nvmlDeviceGetClockInfo(nvmlDevice_t d, nvmlClockType_t t, uint * r) {
    return NVML_nvmlDeviceGetClockInfo(d, t, r);
}

nvmlReturn_t bridge_nvmlDeviceGetPciInfo(nvmlDevice_t d, nvmlPciInfo_t * p) {
    return NVML_nvmlDeviceGetPciInfo(d, p);
}

nvmlReturn_t bridge_nvmlDeviceGetCount(uint * c) {
    return NVML_nvmlDeviceGetCount(c);
}

nvmlReturn_t bridge_nvmlShutdown() {
    if (hDLLN == NULL) return NVML_SUCCESS;
    return NVML_nvmlShutdown();
}

nvmlReturn_t init_nvml() {
    nvmlReturn_t ret;

#ifdef __linux__
    hDLLN = dlopen("libnvidia-ml.so", RTLD_LAZY | RTLD_GLOBAL);
#else
    /* Not in system path, but could be local */
    hDLLN = LoadLibrary("nvml.dll");
    if(!hDLLN) {
        /* %ProgramW6432% is unsupported by OS prior to year 2009 */
        char path[512];
        ExpandEnvironmentStringsA("%ProgramFiles%\\NVIDIA Corporation\\NVSMI\\nvml.dll", path, sizeof(path));
        hDLLN = LoadLibrary(path);
    }
#endif
    if(!hDLLN) {
        PRINTFLN("Unable to load the NVIDIA Management Library");
        return NVML_ERROR_UNINITIALIZED;
    }

    NVML_nvmlInit = (nvmlReturn_t (*)()) dlsym(hDLLN, "nvmlInit_v2");
    if(!NVML_nvmlInit) {
        /* Try an older interface */
        NVML_nvmlInit = (nvmlReturn_t (*)()) dlsym(hDLLN, "nvmlInit");
        if(!NVML_nvmlInit) {
            PRINTFLN("NVML: Unable to initialise");
            return NVML_ERROR_UNINITIALIZED;
        } else {
            NVML_nvmlDeviceGetCount = (nvmlReturn_t (*)(uint *)) \
              dlsym(hDLLN, "nvmlDeviceGetCount");
            NVML_nvmlDeviceGetHandleByIndex = (nvmlReturn_t (*)(uint, nvmlDevice_t *)) \
              dlsym(hDLLN, "nvmlDeviceGetHandleByIndex");
            NVML_nvmlDeviceGetPciInfo = (nvmlReturn_t (*)(nvmlDevice_t, nvmlPciInfo_t *)) \
              dlsym(hDLLN, "nvmlDeviceGetPciInfo");
        }
    } else {
        NVML_nvmlDeviceGetCount = (nvmlReturn_t (*)(uint *)) \
          dlsym(hDLLN, "nvmlDeviceGetCount_v2");
        NVML_nvmlDeviceGetHandleByIndex = (nvmlReturn_t (*)(uint, nvmlDevice_t *)) \
          dlsym(hDLLN, "nvmlDeviceGetHandleByIndex_v2");
        NVML_nvmlDeviceGetPciInfo = (nvmlReturn_t (*)(nvmlDevice_t, nvmlPciInfo_t *)) \
          dlsym(hDLLN, "nvmlDeviceGetPciInfo_v2");
    }

    NVML_nvmlErrorString = (char * (*)(nvmlReturn_t)) \
      dlsym(hDLLN, "nvmlErrorString");
    NVML_nvmlDeviceGetIndex = (nvmlReturn_t (*)(nvmlDevice_t, uint *)) \
      dlsym(hDLLN, "nvmlDeviceGetIndex");
    NVML_nvmlDeviceGetName = (nvmlReturn_t (*)(nvmlDevice_t, char *, uint)) \
      dlsym(hDLLN, "nvmlDeviceGetName");
    NVML_nvmlDeviceGetTemperature = (nvmlReturn_t (*)(nvmlDevice_t, nvmlTemperatureSensors_t, uint *)) \
      dlsym(hDLLN, "nvmlDeviceGetTemperature");
    NVML_nvmlDeviceGetFanSpeed = (nvmlReturn_t (*)(nvmlDevice_t, uint *)) \
      dlsym(hDLLN, "nvmlDeviceGetFanSpeed");
    NVML_nvmlDeviceGetClockInfo = (nvmlReturn_t (*)(nvmlDevice_t, nvmlClockType_t, uint *)) \
      dlsym(hDLLN, "nvmlDeviceGetClockInfo");
    NVML_nvmlDeviceGetUUID = (nvmlReturn_t (*)(nvmlDevice_t, char *, uint)) \
      dlsym(hDLLN, "nvmlDeviceGetUUID");
    NVML_nvmlShutdown = (nvmlReturn_t (*)()) \
      dlsym(hDLLN, "nvmlShutdown");

    return NVML_nvmlInit();
}

int bridge_get_text_property(gettextProperty f,
                             nvmlDevice_t device,
                             char *buf,
                             unsigned int length)
{
    nvmlReturn_t ret;

    ret = f(device, buf, length);

    if (ret == NVML_SUCCESS) {
        return(EXIT_SUCCESS);
    } else {
        return(EXIT_FAILURE);
    }
}

int bridge_get_int_property(getintProperty f,
                             nvmlDevice_t device,
                             unsigned int *property)
{
    nvmlReturn_t ret;

    ret = f(device, property);

    if (ret == NVML_SUCCESS) {
        return(EXIT_SUCCESS);
    } else {
        return(EXIT_FAILURE);
    }
}

#endif
