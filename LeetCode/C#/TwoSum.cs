public class Solution {
    public int[] TwoSum(int[] nums, int target) {
        int[] result = new int[2];
        for(int i=0; i<nums.Length; i++)
        {
            int tempTarget = target;
            if(nums[i] < tempTarget)
            {
                tempTarget -= nums[i];
                for(int j=i+1; j<nums.Length; j++)
                {
                    if(nums[j] == tempTarget)
                    {
                        result[0] = i;
                        result[1] = j;
                        tempTarget = 0;
                        break;
                    }
                }
            }
            if(tempTarget == 0) break;
        }

        return result;
    }
}