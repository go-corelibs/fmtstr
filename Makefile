#!/usr/bin/make --no-print-directory --jobs=1 --environment-overrides -f

VERSION_TAGS += FMTSTR
FMTSTR_MK_SUMMARY := go-corelibs/fmtstr
FMTSTR_MK_VERSION := v1.1.1

include CoreLibs.mk
