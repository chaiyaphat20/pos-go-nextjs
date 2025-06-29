package com.example.app_pos;

import dagger.hilt.InstallIn;
import dagger.hilt.codegen.OriginatingElement;
import dagger.hilt.components.SingletonComponent;
import dagger.hilt.internal.GeneratedEntryPoint;

@OriginatingElement(
    topLevelClass = PosApplication.class
)
@GeneratedEntryPoint
@InstallIn(SingletonComponent.class)
public interface PosApplication_GeneratedInjector {
  void injectPosApplication(PosApplication posApplication);
}
