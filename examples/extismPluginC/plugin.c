#define EXTISM_ENABLE_LOW_LEVEL_API
#define EXTISM_IMPLEMENTATION
#include "./extism-pdk.h"

const char *greeting = "Hello, ";
uint64_t greetingLen = 7;


// TODO: Complete example

int32_t GetDescription() {
  uint64_t inputLen = extism_input_length();

  // Load input
  uint8_t inputData[inputLen];
  extism_load_input(0, inputData, inputLen);

  // Allocate a new offset used to store greeting and name
  uint64_t outputLen = greetingLen + inputLen;
  ExtismPointer offs = extism_alloc(outputLen);
  extism_store(offs, (const uint8_t *)greeting, greetingLen);
  extism_store(offs + greetingLen, inputData, inputLen);

  // Set output
  extism_output_set(offs, outputLen);
  return 0;
}


int32_t test()  {
  uint64_t inputLen = extism_input_length();

  
  // Allocate a new offset used to store greeting and name
  uint64_t outputLen = greetingLen;
  ExtismPointer offs = extism_alloc(outputLen);
  extism_store(offs, (const uint8_t *)greeting, greetingLen);

  // Set output
  extism_output_set(offs, outputLen);
  return 0;
}

int32_t run_test()   {
  uint64_t inputLen = extism_input_length();

  
  // Allocate a new offset used to store greeting and name
  uint64_t outputLen = greetingLen;
  ExtismPointer offs = extism_alloc(outputLen);
  extism_store(offs, (const uint8_t *)greeting, greetingLen);

  // Set output
  extism_output_set(offs, outputLen);
  return 0;
}

