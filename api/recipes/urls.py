from django.urls import path
from recipes import views
from rest_framework.urlpatterns import format_suffix_patterns

urlpatterns = format_suffix_patterns([
    path('', views.api_root),
    path('recipes/', 
        views.RecipeList.as_view(),
        name='recipe-list'
        ),
    path('recipes/<int:pk>',
        views.RecipeDetail.as_view(),
        name='recipe-detail'
        ),
    path('ingredients/',
        views.IngredientList.as_view(),
        name='ingredient-list'
        ),
    path('ingredients/<int:pk>',
        views.IngredientDetail.as_view(),
        name='ingredient-detail'
        )
])