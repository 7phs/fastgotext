#include <stdio.h>
#include "../emd/emd.h"

extern "C" {
    float emd_dist(signature_t *sign1, signature_t *sign2, DistFeatures_t *dist) {
      return emd(sign1, sign2, dist, nullptr, nullptr);
    };
}