package healthService

import (

)

func StartHealthService(kv *kvStore.KVStore) {

}

// should: LOOP: query kv -> heartbeats -> handle accordingly -> update kv + lb