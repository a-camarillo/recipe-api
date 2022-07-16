from django.db import models

# Create your models here.
class Ingredient(models.Model):
    ingredient_name = models.CharField(max_length=100, blank=False)
    recipe = models.ForeignKey(
        'Recipe',
        related_name = 'ingredients',
        on_delete = models.SET_NULL,
        null = True
    )

    def __str__(self):
        return self.ingredient_name

class Recipe(models.Model):
    recipe_name = models.CharField(max_length=200, blank=False) 
    instruction = models.TextField()

    def __str__(self):
        return self.recipe_name

    class Meta:
        ordering = ['recipe_name']