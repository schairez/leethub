<h2><a href="https://leetcode.com/problems/number-of-valid-subarrays/">1063. Number of Valid Subarrays</a></h2><h3>Hard</h3><hr><div><p>Given an integer array <code>nums</code>, return <em>the number of non-empty <strong>subarrays</strong> with the leftmost element of the subarray&nbsp;not larger than other elements in the subarray</em>.</p>

<p>A <strong>subarray</strong> is a <strong>contiguous</strong> part of an array.</p>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>

<pre><strong>Input:</strong> nums = [1,4,2,5,3]
<strong>Output:</strong> 11
<strong>Explanation:</strong> There are 11 valid subarrays: [1],[4],[2],[5],[3],[1,4],[2,5],[1,4,2],[2,5,3],[1,4,2,5],[1,4,2,5,3].
</pre>

<p><strong>Example 2:</strong></p>

<pre><strong>Input:</strong> nums = [3,2,1]
<strong>Output:</strong> 3
<strong>Explanation:</strong> The 3 valid subarrays are: [3],[2],[1].
</pre>

<p><strong>Example 3:</strong></p>

<pre><strong>Input:</strong> nums = [2,2,2]
<strong>Output:</strong> 6
<strong>Explanation:</strong> There are 6 valid subarrays: [2],[2],[2],[2,2],[2,2],[2,2,2].
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 5 * 10<sup>4</sup></code></li>
	<li><code>0 &lt;= nums[i] &lt;= 10<sup>5</sup></code></li>
</ul>
</div>