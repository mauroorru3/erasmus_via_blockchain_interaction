#!/bin/bash

SCRIPTHUB="hub/run-hub.sh"
SCRIPCHAINIT="university_chain_it/run-university_chain_it.sh"
SCRIPCHAINDE="university_chain_de/run-university_chain_de.sh"
SCRIPTTEST="university_chain_it/test.sh"
SCRIPTSTART="test.sh"
SCRIPTRELAYERHERMESCHAINIT="hermes_relayer_it_hub/run-relayer.sh"
SCRIPTRELAYERHERMESCHAINDE="hermes_relayer_de_hub/run-relayer.sh"
SCRIPTRELAYERGOCHAINIT="go_relayer_it_hub/run-relayer_it_hub.sh"
SCRIPTRELAYERGOCHAINDE="go_relayer_de_hub/run-relayer_de_hub.sh"

chmod a+rwx $SCRIPTHUB
chmod a+rwx $SCRIPCHAINIT
chmod a+rwx $SCRIPCHAINDE
chmod a+rwx $SCRIPTTEST
chmod a+rwx $SCRIPTSTART
chmod a+rwx $SCRIPTRELAYERHERMESCHAINIT
chmod a+rwx $SCRIPTRELAYERHERMESCHAINDE
chmod a+rwx $SCRIPTRELAYERGOCHAINIT
chmod a+rwx $SCRIPTRELAYERGOCHAINDE

