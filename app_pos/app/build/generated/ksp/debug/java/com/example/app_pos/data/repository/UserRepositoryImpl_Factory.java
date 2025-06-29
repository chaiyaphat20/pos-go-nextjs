package com.example.app_pos.data.repository;

import com.example.app_pos.data.remote.ApiService;
import dagger.internal.DaggerGenerated;
import dagger.internal.Factory;
import dagger.internal.QualifierMetadata;
import dagger.internal.ScopeMetadata;
import javax.annotation.processing.Generated;
import javax.inject.Provider;

@ScopeMetadata
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
public final class UserRepositoryImpl_Factory implements Factory<UserRepositoryImpl> {
  private final Provider<ApiService> apiServiceProvider;

  public UserRepositoryImpl_Factory(Provider<ApiService> apiServiceProvider) {
    this.apiServiceProvider = apiServiceProvider;
  }

  @Override
  public UserRepositoryImpl get() {
    return newInstance(apiServiceProvider.get());
  }

  public static UserRepositoryImpl_Factory create(Provider<ApiService> apiServiceProvider) {
    return new UserRepositoryImpl_Factory(apiServiceProvider);
  }

  public static UserRepositoryImpl newInstance(ApiService apiService) {
    return new UserRepositoryImpl(apiService);
  }
}
