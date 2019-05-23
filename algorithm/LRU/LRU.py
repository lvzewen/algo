from collections import OrderedDict

class LRUCache:
     """ Implement LRUCache using OrderedDict"""

     def __init__(self, capacity: int):
          self._ordered_dict = OrderedDict()
          self._capacity = capacity

     def get(self, key: int) -> int:
          self._move_to_end_if_exist(key)
          return self._ordered_dict.get(key, -1)

     def put(self, key: int, value: int) -> None:
          self._move_to_end_if_exist(key)

          self._ordered_dict[key] = value
          if len(self._ordered_dict) > self._capacity:
               self._ordered_dict.popitem(last=False)

     def _move_to_end_if_exist(self, key: int) -> int:
          if (key in self._ordered_dict):
               self._ordered_dict.move_to_end(key)