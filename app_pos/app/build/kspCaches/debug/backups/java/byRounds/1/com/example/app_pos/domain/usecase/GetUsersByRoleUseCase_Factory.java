package com.example.app_pos.domain.usecase;

import com.example.app_pos.domain.repository.UserRepository;
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
public final class GetUsersByRoleUseCase_Factory implements Factory<GetUsersByRoleUseCase> {
  private final Provider<UserRepository> repositoryProvider;

  public GetUsersByRoleUseCase_Factory(Provider<UserRepository> repositoryProvider) {
    this.repositoryProvider = repositoryProvider;
  }

  @Override
  public GetUsersByRoleUseCase get() {
    return newInstance(repositoryProvider.get());
  }

  public static GetUsersByRoleUseCase_Factory create(Provider<UserRepository> repositoryProvider) {
    return new GetUsersByRoleUseCase_Factory(repositoryProvider);
  }

  public static GetUsersByRoleUseCase newInstance(UserRepository repository) {
    return new GetUsersByRoleUseCase(repository);
  }
}
