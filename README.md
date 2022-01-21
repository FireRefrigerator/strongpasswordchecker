echo "# strongpasswordchecker" >> README.md
git init
git add README.md
git commit -m "first commit"
git branch -M main
git remote add origin git@github.com:FireRefrigerator/strongpasswordchecker.git
git push -u origin main

1. 算法后期改动较小，更多是替换，所以无需过多抽象，容易理解更重要？算法抽象后效率执行更低。算法重要的是效率（时间复杂度和空间复杂度）