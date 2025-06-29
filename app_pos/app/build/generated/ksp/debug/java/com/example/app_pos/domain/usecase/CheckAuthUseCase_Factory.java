package com.example.app_pos.domain.usecase;

import com.example.app_pos.domain.repository.AuthRepository;
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
public final class CheckAuthUseCase_Factory implements Factory<CheckAuthUseCase> {
  private final Provider<AuthRepository> authRepositoryProvider;

  public CheckAuthUseCase_Factory(Provider<AuthRepository> authRepositoryProvider) {
    this.authRepositoryProvider = authRepositoryProvider;
  }

  @Override
  public CheckAuthUseCase get() {
    return newInstance(authRepositoryProvider.get());
  }

  public static CheckAuthUseCase_Factory create(Provider<AuthRepository> authRepositoryProvider) {
    return new CheckAuthUseCase_Factory(authRepositoryProvider);
  }

  public static CheckAuthUseCase newInstance(AuthRepository authRepository) {
    return new CheckAuthUseCase(authRepository);
  }
}
