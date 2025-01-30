this layer contains development of business logic. A summarized description of the directories in this layer is provided here.
<ul>
   <li>
      Commons directory: Contains repetitive code in this layer. This code is specific to the business logic of this project. In other words, it is not applicable in other projects.
   </li>
   <li>
      Contract directory: Contains interfaces used in this layer. These interfaces are implemented in the Infrastructure layer.
   </li>
   <li>
      Features directory:</strong> Contains the implementation of business logic. Each feature has a handler that is called from REST APIs. The handler consists of four parts:
      <ul>
         <li><strong>Input struct</strong></li>
         <li><strong>Output struct</strong></li>
         <li><strong>Handler struct</strong> (contains the <code>Handle</code> function)</li>
         <li><strong>Validator struct</strong></li>
      </ul>
      The <code>Handle</code> function is invoked from REST APIs.
   </li>
   <li>
      Model directory: Contains DTOs that are only used in this layer. 
   </li>
</ul>
