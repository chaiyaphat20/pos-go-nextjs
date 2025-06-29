package com.example.app_pos.domain.usecase;

import dagger.internal.DaggerGenerated;
import dagger.internal.Factory;
import dagger.internal.QualifierMetadata;
import dagger.internal.ScopeMetadata;
import javax.annotation.processing.Generated;

@ScopeMetadata("javax.inject.Singleton")
@QualifierMetadata
@DaggerGenerated
@Generated(
    value = "dagger.internal.codegen.ComponentProcessor",
    comments = "https://dagger.dev"
)
@SuppressWarnings({
    "unchecked",
    "rawtypes",
    "KotlinInternal",
    "KotlinInternalInJava"
})
public final class CartUseCase_Factory implements Factory<CartUseCase> {
  @Override
  public CartUseCase get() {
    return newInstance();
  }

  public static CartUseCase_Factory create() {
    return InstanceHolder.INSTANCE;
  }

  public static CartUseCase newInstance() {
    return new CartUseCase();
  }

  private static final class InstanceHolder {
    private static final CartUseCase_Factory INSTANCE = new CartUseCase_Factory();
  }
}
