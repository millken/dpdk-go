/*
 * Wraps to support CFunctions with args (...) 
 */

#include <rte_config.h>
#include <rte_common.h>
#include <rte_eal.h>
#include "wrap.h"

inline void rte_eal_exit(int exit_code)
{
    rte_exit(exit_code, "Error with EAL initialization");
}
