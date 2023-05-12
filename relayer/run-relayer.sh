#!/bin/sh

rly config init
rly chains add-dir configs
rly paths add-dir paths

rly keys restore university_chain_it admin_university_chain_it "wrestle rose decorate desert pony tail bulb elder jazz direct law steak cinnamon wife virtual space patient extra segment exile accident flee draw deal"

rly keys restore hub admin_hub "festival vicious churn air remain section exotic quit attack expose host entry raven north repeat silk hire way draft era thumb boil orbit trust"


rly tx link university_chain_it -d -t 3s --src-port hub --dst-port hub --version hub-1
rly start university_chain_it  